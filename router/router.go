package router

import "gopkg.in/macaron.v1"

func SetRouters(m *macaron.Macaron) {
	m.Group("/v1", func() {
		m.Group("/admin", func() {

			m.Group("/organization", func() {

			})

			m.Group("/team", func() {

			})

			m.Group("/project", func() {

			})

			m.Group("/pipeline", func() {

			})

		})

		m.Group("/project", func() {

			// m.Post("/", binding.Bind(handler.ProjectPOSTJSON{}), handler.V1POSTProjectHandler)
			// m.Put("/:project", binding.Bind(handler.ProjectPUTJSON{}), handler.V1PUTProjectHandler)
			// m.Get("/:project", handler.V1GETProjectHandler)
			// m.Delete("/:project", handler.V1DELETEProjectHandler)

		})

		m.Group("/pipeline/:project", func() {

			// m.Post("/", binding.Bind(models.PipelineSpecTemplate{}), handler.V1POSTPipelineHandler)
			// m.Put("/:pipeline", handler.V1PUTPipelineHandler)
			// m.Get("/:pipeline", handler.V1GETPipelineHandler)
			// m.Delete("/:pipeline", handler.V1DELETEPipelineHandler)
			//
			// m.Put("/:pipeline/run", handler.V1RunPipelineHandler)
			//
			// m.Get("/:pipeline/status", handler.V1StatusGETHandler)
		})

		m.Group("/pipelineVersion/:pipeline", func() {
			// m.Post("/", binding.Bind(handler.PointPOSTJSON{}), handler.V1POSTPointHandler)
			// m.Put("/:point", handler.V1PUTPointHandler)
			// m.Get("/:point", handler.V1GETPointHandler)
			// m.Delete("/:point", handler.V1DELETEPointHandler)
		})

		m.Group("/pipelineVersionRun/:pipelineVersion", func() {
			// m.Post("/", binding.Bind(handler.StagePOSTJSON{}), handler.V1POSTStageHandler)
			// m.Put("/:stage", handler.V1PUTStageHandler)
			// m.Get("/:stage", handler.V1GETStageHandler)
			// m.Delete("/:stage", handler.V1DELETEStageHandler)
		})

	})
}
