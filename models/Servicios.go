package models

type servicio struct {
	Id_paquete           int    `json:"id_paquete"`
	Nombre_paquete       string `json:"nombre_paquete"`
	Desc_paquete         string `json:"desc_paquete"`
	Detalle_paquete      string `json:"detalle_paquete"`
	FechaInit            string `json:"fechaInit"`
	FechaFin             string `json:"fechaFin"`
	Cant_dias            int    `json:"cant_dias"`
	Pr_viaje             int    `json:"pr_viaje"`
	Pr_noche             int    `json:"pr_noche"`
	Pr_total             int    `json:"pr_total"`
	Ciudad_origen        string `json:"ciudad_origen"`
	Ciudad_destino       string `json:"ciudad_destino"`
	Nombre_hotel         string `json:"nombre_hotel"`
	Desc_hh              string `json:"desc_hh"`
	Servicios_hh         string `json:"servicios_hh"`
	Rut                  string `json:"rut"`
	Nombre               string `json:"nombre"`
	Apellido             string `json:"apellido"`
	Email                string `json:"email"`
	Fono                 string `json:"fono"`
	Id_servicioAd        int    `json:"Id_servicioAd"`
	Nombre_servicio      string `json:"nombre_servicio"`
	Precio_servicio      int    `json:"precio_servicio"`
	Estado_servicio      string `json:"estado_servicio"`
	Id_factura           int    `json:"id_factura"`
	Fecha_emision        string `json:"fecha_emision"`
	Terminos_condiciones string `json:"terminos_condiciones"`
	Nombre_aerolinea     string `json:"nombre_aerolinea"`
	Metodo_pago          string `json:"metodo_pago"`
	Descuentos           int    `json:"descuentos"`
	Monto_total          int    `json:"monto_total"`
}

var Servicios = []servicio{
	{
		Id_paquete:           1,
		Nombre_paquete:       "Paquete Santiago a Concepción",
		Desc_paquete:         "Viaje desde Santiago a Concepción",
		Detalle_paquete:      "Incluye vuelo y estancia en el Gran Hotel Santiago",
		FechaInit:            "2023-09-05T00:00:00Z",
		FechaFin:             "2023-09-10T00:00:00Z",
		Cant_dias:            5,
		Pr_viaje:             350,
		Pr_noche:             80,
		Pr_total:             750,
		Ciudad_origen:        "Santiago",
		Ciudad_destino:       "Concepción",
		Nombre_hotel:         "Concepción Palace Hotel",
		Desc_hh:              "Habitación individual en el centro de Concepción.",
		Servicios_hh:         "Wi-Fi, TV por cable, Servicio de habitaciones, Aire acondicionado",
		Rut:                  "12345678-9",
		Nombre:               "Juan",
		Apellido:             "Pérez",
		Email:                "juan@example.com",
		Fono:                 "+56912345678",
		Id_servicioAd:        1,
		Nombre_servicio:      "Seguro de Viaje",
		Precio_servicio:      50,
		Estado_servicio:      "Activo",
		Id_factura:           12345,
		Fecha_emision:        "2023-09-15",
		Terminos_condiciones: "Todas las ventas son finales. Política de cancelación: 48 horas antes del vuelo.",
		Nombre_aerolinea:     "Aerolínea Ejemplo",
		Metodo_pago:          "Tarjeta de Crédito",
		Descuentos:           20,
		Monto_total:          750,
	},
}