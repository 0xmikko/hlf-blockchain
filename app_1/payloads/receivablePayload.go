/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package payloads

import (
	"time"
)

type (

	// Data structure with signed order from user
	CreateReceivableReq struct {
		Date   time.Time `json:"date"`
		Issuer int       `json:"issuer"`
		Payer  int       `json:"payer"`
		Amount int       `json:"amount"`
	}
)
