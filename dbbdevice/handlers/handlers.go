package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shiftdevices/godbb/dbbdevice"
	"github.com/shiftdevices/godbb/util/errp"
)

type Handlers struct {
	device dbbdevice.Interface
}

func NewHandlers(
	handleFunc func(string, func(*http.Request) (interface{}, error)) *mux.Route) *Handlers {
	handlers := &Handlers{}

	handleFunc("/status", handlers.getDeviceStatusHandler).Methods("GET")
	handleFunc("/set-password", handlers.postSetPasswordHandler).Methods("POST")
	handleFunc("/create-wallet", handlers.postCreateWalletHandler).Methods("POST")
	handleFunc("/backups/list", handlers.getBackupListHandler).Methods("GET")
	handleFunc("/reset", handlers.postResetDeviceHandler).Methods("POST")
	handleFunc("/login", handlers.postLoginHandler).Methods("POST")
	handleFunc("/backups/erase", handlers.postBackupsEraseHandler).Methods("POST")
	handleFunc("/backups/restore", handlers.postBackupsRestoreHandler).Methods("POST")
	handleFunc("/backups/create", handlers.postBackupsCreateHandler).Methods("POST")

	return handlers
}

func (handlers *Handlers) Init(device dbbdevice.Interface) {
	handlers.device = device
}

func (handlers *Handlers) Uninit() {
	handlers.device = nil
}

func (handlers *Handlers) postSetPasswordHandler(r *http.Request) (interface{}, error) {
	jsonBody := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err != nil {
		return nil, errp.WithStack(err)
	}
	password := jsonBody["password"]
	if err := handlers.device.SetPassword(password); err != nil {
		return map[string]interface{}{"success": false, "errorMessage": err.Error()}, nil
	}
	return map[string]interface{}{"success": true}, nil
}

func (handlers *Handlers) getBackupListHandler(r *http.Request) (interface{}, error) {
	backupList, err := handlers.device.BackupList()
	sdCardInserted := !dbbdevice.IsErrorSDCard(err)
	if sdCardInserted && err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"sdCardInserted": sdCardInserted,
		"backupList":     backupList,
	}, nil
}

func (handlers *Handlers) getDeviceStatusHandler(r *http.Request) (interface{}, error) {
	if handlers.device == nil {
		return "unregistered", nil
	}
	return handlers.device.Status(), nil
}

func (handlers *Handlers) postLoginHandler(r *http.Request) (interface{}, error) {
	jsonBody := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err != nil {
		return nil, errp.WithStack(err)
	}
	password := jsonBody["password"]
	if err := handlers.device.Login(password); err != nil {
		return map[string]interface{}{"success": false, "errorMessage": err.Error()}, nil
	}
	return map[string]interface{}{"success": true}, nil
}

func (handlers *Handlers) postCreateWalletHandler(r *http.Request) (interface{}, error) {
	jsonBody := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err != nil {
		return nil, errp.WithStack(err)
	}
	walletName := jsonBody["walletName"]
	if err := handlers.device.CreateWallet(walletName); err != nil {
		return map[string]interface{}{"success": false, "errorMessage": err.Error()}, nil
	}
	return map[string]interface{}{"success": true}, nil
}

func (handlers *Handlers) postBackupsEraseHandler(r *http.Request) (interface{}, error) {
	jsonBody := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err != nil {
		return nil, errp.WithStack(err)
	}
	filename := jsonBody["filename"]
	return nil, handlers.device.EraseBackup(filename)
}

func (handlers *Handlers) postBackupsRestoreHandler(r *http.Request) (interface{}, error) {
	jsonBody := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err != nil {
		return nil, errp.WithStack(err)
	}
	didRestore, err := handlers.device.RestoreBackup(jsonBody["password"], jsonBody["filename"])
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"didRestore": didRestore}, nil
}

func (handlers *Handlers) postBackupsCreateHandler(r *http.Request) (interface{}, error) {
	jsonBody := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err != nil {
		return nil, errp.WithStack(err)
	}
	return nil, handlers.device.CreateBackup(jsonBody["backupName"])
}

func (handlers *Handlers) postResetDeviceHandler(r *http.Request) (interface{}, error) {
	didReset, err := handlers.device.Reset()
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"didReset": didReset}, nil
}