package jwthelper

func GetBearerToken(header string) string {
	if header != "" {
		return header[len("Bearer "):]
	}

	return ""
}
