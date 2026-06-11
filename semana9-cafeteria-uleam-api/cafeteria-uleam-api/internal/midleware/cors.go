package midleware

import (
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") //acceptar peticiones de cualquier origen
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") //metodos permitidos
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") //encabezados permitidos

		if r.Method == http.MethodOptions { //si la peticion es de tipo OPTIONS
			w.WriteHeader(http.StatusNoContent) //devolver una respuesta vacia
			return 
		}
		
		next.ServeHTTP(w, r) //pasar la peticion al siguiente middleware
	})
}