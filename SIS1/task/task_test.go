package task

import "testing"

// table driven test for creating tasks and checking bad inputs
func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"valid text", "Hello Go world", false},
		{"single word", "Golang", false},
		{"empty string", "", true},
		{"only spaces", "    ", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := NewManager(TextProcessor{})
			created, err := m.CreateTask(tc.input)

			if tc.wantError {
				if err == nil {
					t.Errorf("expected error for input %q, got nil", tc.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if created.ID != 1 {
				t.Errorf("expected ID 1, got %d", created.ID)
			}
			if created.Status != StatusPending {
				t.Errorf("expected status %s, got %s", StatusPending, created.Status)
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	m := NewManager(TextProcessor{})
	created, _ := m.CreateTask("Sample text")

	found, err := m.GetTask(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found.ID != created.ID || found.Input != "Sample text" {
		t.Errorf("got %+v, want %+v", found, created)
	}

	_, err = m.GetTask(999)
	if err == nil {
		t.Errorf("expected error for nonexistent task, got nil")
	}
}

func TestProcessTask(t *testing.T) {
	m := NewManager(TextProcessor{})
	created, _ := m.CreateTask("Hello Go world")

	processed, err := m.ProcessTask(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if processed.Status != StatusCompleted {
		t.Errorf("expected status %s, got %s", StatusCompleted, processed.Status)
	}

	// "Hello Go world" has 3 words, 12 chars without spaces, and 4 vowels (e, o, o, o)
	if processed.Result.WordCount != 3 {
		t.Errorf("expected 3 words, got %d", processed.Result.WordCount)
	}
	if processed.Result.CharCount != 12 {
		t.Errorf("expected 12 chars, got %d", processed.Result.CharCount)
	}
	if processed.Result.VowelCount != 4 {
		t.Errorf("expected 4 vowels, got %d", processed.Result.VowelCount)
	}
}

func TestInvalidStateTransition(t *testing.T) {
	m := NewManager(TextProcessor{})
	created, _ := m.CreateTask("Test transition")

	_, err := m.ProcessTask(created.ID)
	if err != nil {
		t.Fatalf("first process failed: %v", err)
	}

	// try processing again, should fail cause its already completed
	_, err = m.ProcessTask(created.ID)
	if err == nil {
		t.Errorf("expected error when processing completed task again, got nil")
	}
}

func TestMultipleIndependentTasks(t *testing.T) {
	m := NewManager(TextProcessor{})

	t1, _ := m.CreateTask("One two")
	t2, _ := m.CreateTask("One two three four")

	if t1.ID == t2.ID {
		t.Fatalf("tasks should have different IDs, both got %d", t1.ID)
	}

	p1, err := m.ProcessTask(t1.ID)
	if err != nil {
		t.Fatalf("failed to process t1: %v", err)
	}

	// make sure t2 didn't get changed when we processed t1
	checkT2, _ := m.GetTask(t2.ID)
	if checkT2.Status != StatusPending {
		t.Errorf("expected t2 to still be Pending, got %s", checkT2.Status)
	}

	p2, _ := m.ProcessTask(t2.ID)
	if p1.Result.WordCount != 2 {
		t.Errorf("t1 word count = %d, want 2", p1.Result.WordCount)
	}
	if p2.Result.WordCount != 4 {
		t.Errorf("t2 word count = %d, want 4", p2.Result.WordCount)
	}
}
