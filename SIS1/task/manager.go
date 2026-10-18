package task

import "errors"

type Manager struct {
	tasks     map[int]Task
	nextID    int
	processor Processor
}

func NewManager(p Processor) *Manager {
	return &Manager{
		tasks:     make(map[int]Task),
		nextID:    1,
		processor: p,
	}
}

func (m *Manager) CreateTask(input string) (Task, error) {
	t := Task{
		ID:     m.nextID,
		Input:  input,
		Status: StatusPending,
	}

	if err := t.Validate(); err != nil {
		return Task{}, err
	}

	m.tasks[t.ID] = t
	m.nextID++

	return t, nil
}

func (m *Manager) GetTask(id int) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, errors.New("task not found")
	}
	return t, nil
}

func (m *Manager) ProcessTask(id int) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, errors.New("task not found")
	}

	if err := t.UpdateStatus(StatusProcessing); err != nil {
		return Task{}, err
	}
	m.tasks[id] = t

	res, err := m.processor.Process(t.Input)
	if err != nil {
		return Task{}, err
	}

	t.Result = res
	if err := t.UpdateStatus(StatusCompleted); err != nil {
		return Task{}, err
	}

	// need to save it back to the map cause go gives a copy when reading from map[int]Task
	m.tasks[id] = t
	return t, nil
}
