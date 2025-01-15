package main

import "testing"

func validate(address string, expected bool, t *testing.T) {

	label := "valid"
	if !expected {
		label = "invalid"
	}

	t.Logf("Address: %s should be %s", address, label)

	isValidAddress := IsValidAddress(address)

	if isValidAddress != expected {
		t.Log("Failed")
		t.Fail()
		return
	}

	t.Log("Ok")
}

func TestValidAddress(t *testing.T) {
	t.Log("Testing Valid Addresses")

	addresses := []string{
		"TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE",
		"TNaRAoLUyYEV2uF7GUrzSjRQTU8v5ZJ5VR",
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb",
		"TCdb8S5XA6sQYUG3HoHJxVjprKG5KsyKgH",
		"TCf5cqLffPccEY7hcsabiFnMfdipfyryvr",
		"TJDENsfBJs4RFETt1X1W8wMDc8M5XnJhCe",
		"TAzsQ9Gx8eqFNFSKbeXrbi45CuVPHzA8wr",
		"TUL5yxRKeSWvceLZ3BSU5iNJcQmNxkWayh",
		"TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		"TGxy5mJ5325FsVoREcWVkkgHNxo5s5TvFH",
		"TD4VK6d3e3G3jLshNjyBwtw1bVFfxF4LpD",
		"TPBSef2JXqfcXnmJzEX1odCZHQugB5CCZd",
		"TReVPPQjqkDJTyqAH5aZr9DdtgWeZ9UwnJ",
		"TYASr5UV6HEcXatwdFQfmLVUqQQQMUxHLS",
		"TBu1Dk33FQMktAxQviwfVnYSqhtbpqsxHF",
		"TUzqjcFW8L1br3T5ttpUeTZBfsQGYdEGVJ",
		"TXFBqBbqJommqZf7BV8NNYzePh97UmJodJ",
		"TLyqzVGLV1srkB7dToTAEqgDSfPtXRJZYH",
		"TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
		"TYDzsYUEpvnYmQk4zGP9sWWcTEd2MiAtW6",
	}

	for _, item := range addresses {
		validate(item, true, t)
	}
}

func TestInvalidAddress(t *testing.T) {
	t.Log("Testing Invalid Addresses")

	addresses := []string{
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jw",
		"TQrY8bpAXCPYsT7brgEgKfgE4uXwPNYmt",
		"TNaRAoLUyYEV2uF7GUrzSjRQTU8v5ZJ5VRAA",
		"TJDENsfBJs4RFETt1X1W8wMDc8M5XnJhC",
		"TCf5cqLffPccEY7hcsabiFnMfdipfyryv",
		"tYDzsYUEpvnYmQk4zGP9sWWcTEd2MiAtW6",
		"TAzsQ9Gx8eqFNFSKbeXrbi45CuVPHzA8w",
		"TUL5yxRKeSWvceLZ3BSU5iNJcQmNxkWay",
		"TMuA6YqfCeX8EhbfYEg5y7S4DqzSJire",
		"TWsZfpqJHr4TQhQbNbh4hK5h3Rjs8vGZG2BB",
		"1GjYzgCyPobsNS9n6WcbdxKzeh9tQqjvxw",
		"TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLS!",
		"TKE5ENGTvjhzooTQDZkFeVEojF6Qh5Uo8",
		"TNXoiAJ3dct8Fjg4M9fkLFh9S2v8kbPn",
		"TYASr5UV6HEcXatwdFQfmLVUqQQQMUxH",
		"TSaJqQ5N27qYE3ThLosVNT8YrgWWm4hB7",
		"TDtoYxzrKYyJJngXXc3XxqFE4oHYAEhY",
		"0xFBqBbqJommqZf7BV8NNYzePh97UmJodJ",
		"TLyqzVGLV1srkB7dToTAEqgDSfPtXRJZ",
		"TAUN6FwrnwwmaEqYcckffC7wYmbaS6cB",
	}

	for _, item := range addresses {
		validate(item, false, t)
	}
}
