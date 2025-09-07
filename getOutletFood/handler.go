package getOutletFood

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gorm.io/datatypes"
)

// in-memory holder kept as fallback
var outletFoodData Response

// helper to write JSON
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("json encode error:", err)
	}
}

// getOutletFoodHandler:
// 1) Validate request body (same as before).
// 2) Try DB: fetch latest OutletFood row for req.OutletID and return it if found.
// 3) Else if mock_response.json exists return it.
// 4) Else fallback to in-memory mock.
func getOutletFoodHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// basic validation
	if req.Platform == "" || req.Country == "" || req.City == "" || req.OutletID == 0 || req.Date == "" {
		http.Error(w, "missing required field (platform,country,city,outletid,date)", http.StatusBadRequest)
		return
	}

	// parse date (expects RFC3339)
	if _, err := time.Parse(time.RFC3339, req.Date); err != nil {
		http.Error(w, "invalid date format, expect RFC3339: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 1) Try DB if connected and OutletID provided
	if DB != nil && req.OutletID != 0 {
		var rec OutletFood
		// latest record for this outlet_id
		if err := DB.Where("outlet_id = ?", req.OutletID).Order("updated_at desc").First(&rec).Error; err == nil {
			var resp Response
			if err := json.Unmarshal(rec.Data, &resp); err == nil {
				writeJSON(w, resp)
				return
			}
			log.Println("db json unmarshal error:", err)
			// fallthrough to other options if DB row corrupt
		}
	}

	// 2) Option: if you have mock_response.json (copied from PDF), return it exactly:
	if data, err := os.ReadFile("mock_response.json"); err == nil {
		var resp Response
		if err := json.Unmarshal(data, &resp); err == nil {
			writeJSON(w, resp)
			return
		}
		log.Println("mock_response.json exists but unmarshal failed:", err)
	}

	// 3) Fall back to in-memory mock
	if outletFoodData.Status == 0 {
		outletFoodData = Response{
			Status: 200,
			Code:   10001,
			Result: "success",
			Msg:    "",
			Output: Output{
				OutletName: "Default INOX",
				City:       CityInfo{Name: "Default City", State: "Default State"},
				R: []Item{
					{ID: "I0001", H: "Sample Popcorn", Dp: 50000, Ct: "POPCORN", Veg: true, Wt: "100 g", En: "300 kcal"},
				},
				Cat: []string{"POPCORN"},
			},
		}
	}

	writeJSON(w, outletFoodData)
}

// updateOutletFoodHandler:
// Accepts the full Response JSON body (same shape as your PDF). It will detect an outletId
// either from query param `?outletid=200` or from newData.Output.City.ID (if present).
// Then it stores the full JSON into Postgres JSONB (upsert semantics: create or update).
func updateOutletFoodHandler(w http.ResponseWriter, r *http.Request) {
	var newData Response
	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// determine outletID: priority -> query param `outletid`, else newData.Output.City.ID
	outletID := 0
	if q := r.URL.Query().Get("outletid"); q != "" {
		if v, err := strconv.Atoi(q); err == nil {
			outletID = v
		}
	}
	if outletID == 0 && newData.Output.City.ID != 0 {
		outletID = newData.Output.City.ID
	}

	// If still missing, return error (you can choose to allow outletID==0 if you prefer)
	if outletID == 0 {
		http.Error(w, "missing outletid (send as ?outletid=xxx or include output.city.id)", http.StatusBadRequest)
		return
	}

	// marshal the Response to bytes to store as JSONB
	b, err := json.Marshal(newData)
	if err != nil {
		http.Error(w, "failed to marshal response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// persist in DB if available
	if DB != nil {
		var rec OutletFood
		if err := DB.Where("outlet_id = ?", outletID).First(&rec).Error; err == nil {
			// exists -> update
			rec.Data = datatypes.JSON(b)
			if err := DB.Save(&rec).Error; err != nil {
				http.Error(w, "db save error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			// create new
			rec = OutletFood{
				OutletID: outletID,
				Data:     datatypes.JSON(b),
			}
			if err := DB.Create(&rec).Error; err != nil {
				http.Error(w, "db create error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	// update in-memory and persist to file (optional, keeps backwards compatibility)
	outletFoodData = newData
	if data, err := json.MarshalIndent(newData, "", "  "); err == nil {
		_ = os.WriteFile("mock_response.json", data, 0644)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Data updated successfully"})
}
