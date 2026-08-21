package tracing

import (
	"testing"
	"time"
)

func TestCollectorRecordsNestedSpans(t *testing.T) {
	collector := NewCollector()

	collector.Trace("root", func(root *Span) {
		root.Trace("child", func(child *Span) {
			child.Trace("grandchild", func(*Span) {})
		})
	})

	spans := collector.Spans()
	if len(spans) != 1 {
		t.Fatalf("expected 1 root span, got %d", len(spans))
	}
	if spans[0].Name != "root" {
		t.Errorf("expected root name %q, got %q", "root", spans[0].Name)
	}
	if len(spans[0].Children) != 1 {
		t.Fatalf("expected 1 child span, got %d", len(spans[0].Children))
	}
	if spans[0].Children[0].Name != "child" {
		t.Errorf("expected child name %q, got %q", "child", spans[0].Children[0].Name)
	}
	if len(spans[0].Children[0].Children) != 1 {
		t.Fatalf("expected 1 grandchild span, got %d", len(spans[0].Children[0].Children))
	}
	if spans[0].Children[0].Children[0].Name != "grandchild" {
		t.Errorf("expected grandchild name %q, got %q", "grandchild", spans[0].Children[0].Children[0].Name)
	}
}

func TestCollectorRecordsDuration(t *testing.T) {
	collector := NewCollector()

	collector.Trace("root", func(*Span) {
		time.Sleep(time.Millisecond)
	})

	if duration := collector.Spans()[0].Duration; duration < time.Millisecond {
		t.Errorf("expected duration of at least %s, got %s", time.Millisecond, duration)
	}
}
