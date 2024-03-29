package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var out In
	for _, stage := range stages {
		out = func(in In, done In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for {
					select {
					case <-done:
						return
					case val, ok := <-in:
						if !ok {
							return
						}
						select {
						case out <- val:
						case <-done:
							return
						}
					}
				}
			}()
			return out
		}(in, done)
		in = stage(out)
	}
	return in
}
