package sectorstorage

import "github.com/kilb/sector-storage/storiface"

func (m *Manager) WorkerStats() map[uint64]storiface.WorkerStats {
	m.sched.workersLk.Lock()
	defer m.sched.workersLk.Unlock()

	out := map[uint64]storiface.WorkerStats{}

	for id, handle := range m.sched.workers {
		out[uint64(id)] = storiface.WorkerStats{
			Info:       handle.info,
			MemUsedMin: handle.active.memUsedMin,
			MemUsedMax: handle.active.memUsedMax,
			GpuUsed:    handle.active.gpuUsed,
			CpuUse:     handle.active.cpuUse,
		}
	}

	return out
}
