package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// prodi struct (model)
type Prodis struct {
	ProdiID          string `json:"ProdiID"`
	KodeID           string `json:"KodeID"`
	FakultasID       string `json:"FakultasID"`
	Nama             string `json:"Nama"`
	Kodevir          string `json:"kodevir"`
	KodeExplode      string `json:"KodeExplode"`
	NoRekening       string `json:"NoRekening"`
	Nama_en          string `json:"Nama_en"`
	SINGKATAN        string `json:"SINGKATAN"`
	JenjangID        string `json:"JenjangID"`
	Gelar            string `json:"Gelar"`
	ProdiDiktiID     string `json:"ProdiDiktiID"`
	NamaSesi         string `json:"NamaSesi"`
	Akreditasi       string `json:"Akreditasi"`
	NoSKDikti        string `json:"NoSKDikti"`
	TglSKDikti       string `json:"tglSKDikti"`
	NoSKBAN          string `json:"NoSKBAN"`
	TglSKBAN         string `json:"TglSKBAN"`
	PajakHonorDosen  string `json:"PajakHonorDosen"`
	Pejabat          string `json:"Pejabat"`
	Jabatan          string `json:"Jabatan"`
	FotmatNim        string `json:"FotmatNim"`
	DapatPindahProdi string `json:"DapatPindahProdi"`
	DefSKS           string `json:"DefSKS"`
	TotalSKS         string `json:"TotalSKS"`
	DefKehadiran     string `json:"DefKehadiran"`
	BatasStudi       string `json:"atasStudi"`
	JumlahSesi       string `json:"JumlahSesi"`
	CekPrasyarat     string `json:"CekPrasyarat"`
	LoginBuat        string `json:"LoginBuat"`
	TanggalBuat      string `json:"TanggalBuat"`
	LoginEdit        string `json:"LoginEdit"`
	TanggalEdit      string `json:"TanggalEdit"`
	Keterangan       string `json:"Keterangan"`
	StartNoProdi     string `json:"StartNoProdi"`
	NoProdi          string `json:"NoProdi"`
	Denda1           string `json:"Denda1"`
	Denda2           string `json:"Denda2"`
	NA               string `json:"NA"`
}

// get all orders
func getProdies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var prodi []Prodis

	sql := `SELECT
				ProdiID,
				IFNULL(KodeID,''),
				IFNULL(FakultasID,'') FakultasID,
				IFNULL(Nama,'') Nama,
				IFNULL(kodevir,'') kodevir,
				IFNULL(KodeExplode,'') KodeExplode,
				IFNULL(NoRekening,'') NoRekening,
				IFNULL(Nama_en,'') Nama_en,
				IFNULL(SINGKATAN,'') SINGKATAN,
				IFNULL(JenjangID,'') JenjangID,
				IFNULL(Gelar,'') Gelar,
				IFNULL(ProdiDiktiID,'') ProdiDiktiID,
				IFNULL(NamaSesi,'') NamaSesi,
				IFNULL(Akreditasi,'') Akreditasi,
				IFNULL(NoSKDikti,'') NoSKDikti,
				IFNULL(NOSKBAN,'') NOSKBAN,
				IFNULL(TglSKBAN,'') TglSKBAN,
				IFNULL(PajakHonorDosen,'') PajakHonorDosen,
				IFNULL(Pejabat,'') Pejabat,
				IFNULL(FormatNim,'')FormatNim,
				IFNULL(DapatPindahProdi,'') DapatPindahProdi,
				IFNULL(DefSKS,'') DefSKS,
				IFNULL(TotalSKS,'') TotalSKS,
				IFNULL(DefKehadiran,'') DefKehadiran,
				IFNULL(BatasStudi,'') BatasStudi,
				IFNULL(JumlahSesi,'') JumlahSesi,
				IFNULL(CekPrasyarat,'') CekPrasyarat,
				IFNULL(LoginBuat,'') LoginBuat,
				IFNULL(TanggalBuat,'') TanggalBuat,
				IFNULL(LoginEdit,'') LoginEdit,
				IFNULL(TanggalEdit,'') TanggalEdit,
				IFNULL(Keterangan,'') Keterangan,
				IFNULL(StartNoProdi,'') StartNoProdi,
				IFNULL(NoProdi,'') NoProdi,
				IFNULL(Denda1,'') Denda1,
				IFNULL(Denda2,'') Denda2,
				IFNULL(NA,'') NA
			FROM prodi`
	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}
	for result.Next() {

		var prodis Prodis
		err := result.Scan(&prodis.ProdiID, &prodis.KodeID, &prodis.FakultasID, &prodis.Nama, &prodis.Kodevir, &prodis.KodeExplode,
			&prodis.NoRekening, &prodis.Nama_en, &prodis.SINGKATAN, &prodis.JenjangID, &prodis.Gelar, &prodis.ProdiDiktiID, &prodis.NamaSesi,
			&prodis.Akreditasi, &prodis.NoSKDikti, &prodis.TglSKDikti, &prodis.NoSKBAN, &prodis.TglSKBAN, &prodis.PajakHonorDosen, &prodis.Pejabat,
			&prodis.Jabatan, &prodis.FotmatNim, &prodis.DapatPindahProdi, &prodis.DefSKS, &prodis.TotalSKS, &prodis.DefKehadiran, &prodis.BatasStudi,
			&prodis.JumlahSesi, &prodis.CekPrasyarat, &prodis.LoginBuat, &prodis.TanggalBuat, &prodis.LoginEdit, &prodis.TanggalEdit, &prodis.Keterangan,
			&prodis.StartNoProdi, &prodis.NoProdi, &prodis.Denda1, &prodis.Denda2, &prodis.NA)

		if err != nil {
			panic(err.Error())
		}
		prodi = append(prodi, prodis)
	}
	json.NewEncoder(w).Encode(prodi)
}

func createProdi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		ProdiID := r.FormValue("ProdiID")
		KodeID := r.FormValue("KodeID")
		FakultasID := r.FormValue("FakultasID")
		Nama := r.FormValue("Nama")
		kodevir := r.FormValue("kodevir")
		KodeExplode := r.FormValue("KodeExplode")
		NoRekening := r.FormValue("NoRekening")
		Nama_en := r.FormValue("Nama_en")
		SINGKATAN := r.FormValue("SINGKATAN")
		JenjangID := r.FormValue("JenjangID")
		Gelar := r.FormValue("Gelar")
		ProdiDiktiID := r.FormValue("ProdiDiktiID")
		NamaSesi := r.FormValue("NamaSesi")
		Akreditasi := r.FormValue("Akreditasi")
		NoSKDikti := r.FormValue("NoSKDikti")
		TglSKDikti := r.FormValue("TglSKDikti")
		NoSKBAN := r.FormValue("NoSKBAN")
		TglSKBAN := r.FormValue("TglSKBAN")
		PajakHonorDosen := r.FormValue("PajakHonorDosen")
		Pejabat := r.FormValue("Pejabat")
		Jabatan := r.FormValue("Jabatan")
		FormatNim := r.FormValue("FormatNim")
		DapatPindahProdi := r.FormValue("DapatPindahProdi")
		DefSKS := r.FormValue("DefSKS")
		TotalSKS := r.FormValue("TotalSKS")
		DefKehadiran := r.FormValue("DefKehadiran")
		BatasStudi := r.FormValue("BatasStudi")
		JumlahSesi := r.FormValue("JumlahSesi")
		CekPrasyarat := r.FormValue("CekPrasyarat")
		LoginBuat := r.FormValue("LoginBuat")
		TanggalBuat := r.FormValue("TanggalBuat")
		LoginEdit := r.FormValue("LoginEdit")
		TanggalEdit := r.FormValue("TanggalEdit")
		Keterangan := r.FormValue("Keterangan")
		StartNoProdi := r.FormValue("StartNoProdi")
		NoProdi := r.FormValue("NoProdi")
		Denda1 := r.FormValue("Denda1")
		Denda2 := r.FormValue("Denda2")
		NA := r.FormValue("NA")

		stmt, err := db.Prepare("INSERT INTO prodi (ProdiID,KodeID,FakultasID,Nama,kodevir,KodeExplode,NoRekening,Nama_en,SINGKATAN,JenjangID,Gelar,ProdiDiktiID,NamaSesi,Akreditasi,NoSKDikti,TglSKDikti,NoSKBAN,TglSKBAN,PajakHonorDosen,Pejabat,Jabatan,FormatNim,DapatPindahProdi,DefSKS,Totalsks,DefKehadiran,BatasStudi,JumlahSesi,CekPrasyarat,LoginBuat,TanggalBuat,LoginEdit,TanggalEdit,Keterangan,StartNoProdi,NoProdi,Denda1,Denda2,NA) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(ProdiID, KodeID, FakultasID, Nama, kodevir, KodeExplode, NoRekening, Nama_en,
			SINGKATAN, JenjangID, Gelar, ProdiDiktiID, NamaSesi, Akreditasi, NoSKDikti, TglSKDikti, NoSKBAN,
			TglSKBAN, PajakHonorDosen, Pejabat, Jabatan, FormatNim, DapatPindahProdi, DefSKS, TotalSKS, DefKehadiran,
			BatasStudi, JumlahSesi, CekPrasyarat, LoginBuat, TanggalBuat, LoginEdit, TanggalEdit, Keterangan, StartNoProdi, NoProdi,
			Denda1, Denda2, NA)
		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}
func getProdi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prodi []Prodis
	params := mux.Vars(r)

	sql := `SELECT
				ProdiID,
				IFNULL(KodeID,''),
				IFNULL(FakultasID,'') FakultasID,
				IFNULL(Nama,'') Nama,
				IFNULL(kodevir,'') kodevir,
				IFNULL(KodeExplode,'') KodeExplode,
				IFNULL(NoRekening,'') NoRekening,
				IFNULL(Nama_en,'') Nama_en,
				IFNULL(SINGKATAN,'') SINGKATAN,
				IFNULL(JenjangID,'') JenjangID,
				IFNULL(Gelar,'') Gelar,
				IFNULL(ProdiDiktiID,'') ProdiDiktiID,
				IFNULL(NamaSesi,'') NamaSesi,
				IFNULL(Akreditasi,'') Akreditasi,
				IFNULL(NoSKDikti,'') NoSKDikti,
				IFNULL(NOSKBAN,'') NOSKBAN,
				IFNULL(TglSKBAN,'') TglSKBAN,
				IFNULL(PajakHonorDosen,'') PajakHonorDosen,
				IFNULL(Pejabat,'') Pejabat,
				IFNULL(FormatNim,'')FormatNim,
				IFNULL(DapatPindahProdi,'') DapatPindahProdi,
				IFNULL(DefSKS,'') DefSKS,
				IFNULL(TotalSKS,'') TotalSKS,
				IFNULL(DefKehadiran,'') DefKehadiran,
				IFNULL(BatasStudi,'') BatasStudi,
				IFNULL(JumlahSesi,'') JumlahSesi,
				IFNULL(CekPrasyarat,'') CekPrasyarat,
				IFNULL(LoginBuat,'') LoginBuat,
				IFNULL(TanggalBuat,'') TanggalBuat,
				IFNULL(LoginEdit,'') LoginEdit,
				IFNULL(TanggalEdit,'') TanggalEdit,
				IFNULL(Keterangan,'') Keterangan,
				IFNULL(StartNoProdi,'') StartNoProdi,
				IFNULL(NoProdi,'') NoProdi,
				IFNULL(Denda1,'') Denda1,
				IFNULL(Denda2,'') Denda2,
				IFNULL(NA,'') NA
			FROM prodi WHERE ProdiID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var prodis Prodis
	for result.Next() {
		err := result.Scan(&prodis.ProdiID, &prodis.KodeID, &prodis.FakultasID, &prodis.Nama, &prodis.Kodevir, &prodis.KodeExplode,
			&prodis.NoRekening, &prodis.Nama_en, &prodis.SINGKATAN, &prodis.JenjangID, &prodis.Gelar, &prodis.ProdiDiktiID, &prodis.NamaSesi,
			&prodis.Akreditasi, &prodis.NoSKDikti, &prodis.TglSKDikti, &prodis.NoSKBAN, &prodis.TglSKBAN, &prodis.PajakHonorDosen, &prodis.Pejabat,
			&prodis.Jabatan, &prodis.FotmatNim, &prodis.DapatPindahProdi, &prodis.DefSKS, &prodis.TotalSKS, &prodis.DefKehadiran, &prodis.BatasStudi,
			&prodis.JumlahSesi, &prodis.CekPrasyarat, &prodis.LoginBuat, &prodis.TanggalBuat, &prodis.LoginEdit, &prodis.TanggalEdit, &prodis.Keterangan,
			&prodis.StartNoProdi, &prodis.NoProdi, &prodis.Denda1, &prodis.Denda2, &prodis.NA)

		if err != nil {
			panic(err.Error())
		}
		prodi = append(prodi, prodis)
	}
	json.NewEncoder(w).Encode(prodi)
}
func updateProdi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		params := mux.Vars(r)
		newKodeID := r.FormValue("KodeID")
		stmt, err := db.Prepare("UPDATE prodi SET KodeID = ? WHERE ProdiID = ?")
		_, err = stmt.Exec(newKodeID, params["id"])
		if err != nil {
			fmt.Fprintf(w, " data not found or request error")
		}
		fmt.Fprintf(w, "Kode with ID = %s was deleted", params["id"])
	}
}

func deleteProdi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM prodi WHERE ProdiID = ?")

	_, err = stmt.Exec(params["id"])
	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}
	fmt.Fprintf(w, "Prodi with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var prodi []Prodis

	ProdiID := r.FormValue("ProdiID")
	KodeID := r.FormValue("KodeID")

	sql := `SELECT
				ProdiID,
				IFNULL(KodeID,''),
				IFNULL(FakultasID,'') FakultasID,
				IFNULL(Nama,'') Nama,
				IFNULL(kodevir,'') kodevir,
				IFNULL(KodeExplode,'') KodeExplode,
				IFNULL(NoRekening,'') NoRekening,
				IFNULL(Nama_en,'') Nama_en,
				IFNULL(SINGKATAN,'') SINGKATAN,
				IFNULL(JenjangID,'') JenjangID,
				IFNULL(Gelar,'') Gelar,
				IFNULL(ProdiDiktiID,'') ProdiDiktiID,
				IFNULL(NamaSesi,'') NamaSesi,
				IFNULL(Akreditasi,'') Akreditasi,
				IFNULL(NoSKDikti,'') NoSKDikti,
				IFNULL(NOSKBAN,'') NOSKBAN,
				IFNULL(TglSKBAN,'') TglSKBAN,
				IFNULL(PajakHonorDosen,'') PajakHonorDosen,
				IFNULL(Pejabat,'') Pejabat,
				IFNULL(FormatNim,'')FormatNim,
				IFNULL(DapatPindahProdi,'') DapatPindahProdi,
				IFNULL(DefSKS,'') DefSKS,
				IFNULL(TotalSKS,'') TotalSKS,
				IFNULL(DefKehadiran,'') DefKehadiran,
				IFNULL(BatasStud,''i) BatasStudi,
				IFNULL(JumlahSesi,'') JumlahSesi,
				IFNULL(CekPrasyarat,'') CekPrasyarat,
				IFNULL(LoginBuat,'') LoginBuat,
				IFNULL(TanggalBuat,'') TanggalBuat,
				IFNULL(LoginEdit,'') LoginEdit,
				IFNULL(TanggalEdit,'') TanggalEdit,
				IFNULL(Keterangan,'') Keterangan,
				IFNULL(StartNoProdi,'') StartNoProdi,
				IFNULL(NoProdi,'') NoProdi,
				IFNULL(Denda1,'') Denda1,
				IFNULL(Denda2,'') Denda2,
				IFNULL(NA,'') NA
			FROM prodi WHERE ProdiID = ? AND KodeID = ?`
	result, err := db.Query(sql, ProdiID, KodeID)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var prodis Prodis

	for result.Next() {
		err := result.Scan(&prodis.ProdiID, &prodis.KodeID, &prodis.FakultasID, &prodis.Nama, &prodis.Kodevir, &prodis.KodeExplode,
			&prodis.NoRekening, &prodis.Nama_en, &prodis.SINGKATAN, &prodis.JenjangID, &prodis.Gelar, &prodis.ProdiDiktiID, &prodis.NamaSesi,
			&prodis.Akreditasi, &prodis.NoSKDikti, &prodis.TglSKDikti, &prodis.NoSKBAN, &prodis.TglSKBAN, &prodis.PajakHonorDosen, &prodis.Pejabat,
			&prodis.Jabatan, &prodis.FotmatNim, &prodis.DapatPindahProdi, &prodis.DefSKS, &prodis.TotalSKS, &prodis.DefKehadiran, &prodis.BatasStudi,
			&prodis.JumlahSesi, &prodis.CekPrasyarat, &prodis.LoginBuat, &prodis.TanggalBuat, &prodis.LoginEdit, &prodis.TanggalEdit, &prodis.Keterangan,
			&prodis.StartNoProdi, &prodis.NoProdi, &prodis.Denda1, &prodis.Denda2, &prodis.NA)
		if err != nil {
			panic(err.Error())
		}
		prodi = append(prodi, prodis)
	}
	json.NewEncoder(w).Encode(prodi)
}

//main function
func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//init router
	r := mux.NewRouter()

	//route hendles
	r.HandleFunc("/prodi", getProdies).Methods("GET")
	r.HandleFunc("/prodi/{id}", getProdi).Methods("GET")
	r.HandleFunc("/prodi", createProdi).Methods("POST")
	r.HandleFunc("/prodi/{id}", updateProdi).Methods("PUT")
	r.HandleFunc("/prodi/{id}", deleteProdi).Methods("DELETE")

	//new
	r.HandleFunc("/getprodis", getPost).Methods("POST")

	//start server
	log.Fatal(http.ListenAndServe(":8282", r))
}
