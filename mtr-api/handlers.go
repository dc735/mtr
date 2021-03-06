package main

import (
	"bytes"
	"github.com/GeoNet/weft"
	"net/http"
)

func appHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var a appID

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return a.proto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func appMetricHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var a appMetric

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		default:
			return a.svg(r, h, b)
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func applicationMetricHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var a applicationMetric

	switch r.Method {
	case "PUT":
		return a.put(r)
	default:
		return &weft.MethodNotAllowed
	}
}

func applicationCounterHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var a applicationCounter

	switch r.Method {
	case "PUT":
		return a.put(r)
	default:
		return &weft.MethodNotAllowed
	}
}

func applicationTimerHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var a applicationTimer

	switch r.Method {
	case "PUT":
		return a.put(r)
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldMetricHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldMetric

	switch r.Method {
	case "PUT":
		return f.put(r)
	case "DELETE":
		return f.delete(r)
	case "GET":
		switch r.Header.Get("Accept") {
		default:
			return f.svg(r, h, b)
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldMetricTagHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldMetricTag

	switch r.Method {
	case "PUT":
		return f.put(r, h, b)
	case "DELETE":
		return f.delete(r, h, b)
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.all(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func tagHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var t tag
	var ts tagSearch

	switch r.Method {
	case "PUT":
		return t.put(r)
	case "DELETE":
		return t.delete(r)
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return ts.singleProto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func tagsHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var ts tagSearch

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return ts.allProto(r, h, b)
		default:
			return &weft.NotAcceptable
		}

	default:
		return &weft.MethodNotAllowed
	}
}

func fieldModelHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldModel

	switch r.Method {
	case "PUT":
		return f.put(r)
	case "DELETE":
		return f.delete(r)
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.proto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldDeviceHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldDevice

	switch r.Method {
	case "PUT":
		return f.put(r)
	case "DELETE":
		return f.delete(r)
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.proto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldTypeHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldType

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.proto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldThresholdHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldThreshold

	switch r.Method {
	case "PUT":
		return f.save(r)
	case "DELETE":
		return f.delete(r)
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.proto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldMetricLatestHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldLatest

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.proto(r, h, b)
		case "application/vnd.geo+json":
			return f.geoJSON(r, h, b)
		default:
			return f.svg(r, h, b)
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldStateHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldState

	switch r.Method {
	case "PUT":
		return f.save(r)
	case "DELETE":
		return f.delete(r)
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.allProto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func fieldStateTagHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f fieldStateTag

	switch r.Method {
	case "PUT":
		return f.put(r, h, b)
	case "DELETE":
		return f.delete(r, h, b)
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.all(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func dataSiteHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var d dataSite

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return d.allProto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	case "PUT":
		return d.put(r)
	case "DELETE":
		return d.delete(r)
	default:
		return &weft.MethodNotAllowed
	}
}

func dataLatencyHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var d dataLatency

	switch r.Method {
	case "PUT":
		return d.put(r)
	case "DELETE":
		return d.delete(r)
	case "GET":
		switch r.Header.Get("Accept") {
		default:
			return d.svg(r, h, b)
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func dataLatencyThresholdHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var d dataLatencyThreshold

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return d.get(r, h, b)
		default:
			return &weft.NotAcceptable
		}

	case "PUT":
		return d.save(r)
	case "DELETE":
		return d.delete(r)
	default:
		return &weft.MethodNotAllowed
	}
}

func dataLatencyTagHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f dataLatencyTag

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.all(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	case "PUT":
		return f.put(r, h, b)
	case "DELETE":
		return f.delete(r, h, b)
	default:
		return &weft.MethodNotAllowed
	}
}

func dataLatencySummaryHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var d dataLatencySummary

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return d.proto(r, h, b)
		default:
			return d.svg(r, h, b)
		}
	default:
		return &weft.MethodNotAllowed
	}
}

func dataTypeHandler(r *http.Request, h http.Header, b *bytes.Buffer) *weft.Result {
	var f dataType

	switch r.Method {
	case "GET":
		switch r.Header.Get("Accept") {
		case "application/x-protobuf":
			return f.proto(r, h, b)
		default:
			return &weft.NotAcceptable
		}
	default:
		return &weft.MethodNotAllowed
	}
}
