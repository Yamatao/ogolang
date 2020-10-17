package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageWrapper(outerIn In, stageIn chan interface{}, stageOut Out, outerOut chan interface{}, done In) {
	defer close(outerOut)
	defer close(stageIn)
	for {
		select {
		// wait for something from outer scope
		case i, ok := <-outerIn:
			if !ok {
				return
			}
			// now push the `i` to a real stage
			select {
			case stageIn <- i:
				break
			case <-done:
				return
			}
			// and wait a result from the stage
			select {
			case o := <-stageOut:
				outerOut <- o
			case <-done:
				return
			}
		case <-done:
			return
		}
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) (out Out) {
	ch := in
	for _, stage := range stages {
		stageIn := make(chan interface{}) // auxiliary channels for the wrapper
		outerOut := make(chan interface{})

		stageOut := stage(stageIn)
		go stageWrapper(ch, stageIn, stageOut, outerOut, done)
		ch = outerOut
	}
	out = ch
	return
}
