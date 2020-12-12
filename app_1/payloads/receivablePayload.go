/*
 * Copyright (c) 2020. DLT Experts.
 *  Authors: Mikael Lazarev, Ivan Fedorov
 */

package payloads

type (

	// Data structure with signed order from user
	CreateReceivableReq struct {
		Issuer string `json:"issuer"`
		Payer  string `json:"payer"`
		Amount int    `json:"amount"`
	}
)
