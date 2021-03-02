package router

import (
	"net/http"
	"simpus/app/keuangan"
	"simpus/app/pelayanan/pendaftaran/antrian"
	"simpus/app/pelayanan/pendaftaran/pasien"
	"simpus/app/settings/config"

	"github.com/gorilla/mux"
)

// Init is
func Init() {
	r := mux.NewRouter()

	// Pelayanan > Pendaftaran > Pasien

	r.HandleFunc("/api/pelayanan/pendaftaran/pasien", pasien.Index).Methods("GET")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien", pasien.Store).Methods("POST")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien/{id}", pasien.Show).Methods("GET")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien/{id}", pasien.Update).Methods("PUT")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasien/{id}", pasien.Destroy).Methods("DELETE")
	r.HandleFunc("/api/pelayanan/pendaftaran/pasienCount", pasien.GetCountPasien).Methods("GET")

	// Pelayanan > Pendaftaran > Antrian

	r.HandleFunc("/api/pelayanan/pendaftaran/antrian", antrian.Index).Methods("GET")
	r.HandleFunc("/api/pelayanan/pendaftaran/antrian", antrian.Store).Methods("POST")

	// Keuangan > COA

	r.HandleFunc("/api/coa", keuangan.IndexCOA).Methods("GET")
	r.HandleFunc("/api/coa", keuangan.StoreCOA).Methods("POST")
	r.HandleFunc("/api/coa/{id}", keuangan.ShowCOA).Methods("GET")
	r.HandleFunc("/api/coa/{id}", keuangan.UpdateCOA).Methods("PUT")
	r.HandleFunc("/api/coa/{id}", keuangan.DestroyCOA).Methods("DELETE")
	r.HandleFunc("/api/coa/{id}/subaccount", keuangan.StoreSubAccountCOA).Methods("POST")
	r.HandleFunc("/api/coa/{id}/subaccount/{index}", keuangan.DestroySubAccountCOA).Methods("DELETE")

	r.HandleFunc("/api/db/users", config.GetUser).Methods("GET")
	r.HandleFunc("/api/db", config.IndexDB).Methods("GET")
	r.HandleFunc("/api/db/collections", config.IndexCollection).Methods("GET")
	// r.HandleFunc("/api/db/collection", config.Store).Methods("POST")

	http.ListenAndServe(":1234", r)
}