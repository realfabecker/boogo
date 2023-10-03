package spin

import (
	"time"

	"github.com/pterm/pterm"
)

type spinner struct {
	endTime      time.Time
	startTime    time.Time
	ptermSpinner *pterm.SpinnerPrinter
}

func NewSpinner(text string) *spinner {
	return &spinner{
		startTime:    time.Now(),
		ptermSpinner: pterm.DefaultSpinner.WithText(text),
	}
}

func (sp *spinner) WrapStart(fn func() error) error {
	sp.ptermSpinner.Start()
	if err := fn(); err != nil {
		sp.endTime = time.Now()
		return err
	}
	sp.endTime = time.Now()
	sp.ptermSpinner.Success()
	return nil
}
