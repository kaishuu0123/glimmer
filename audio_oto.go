//go:build !windows
// +build !windows

package glimmer

import "github.com/hajimehoshi/oto"

type audioOutput struct {
	ab *AudioBuffer

	ctx    *oto.Context
	player *oto.Player
}

func (ao *audioOutput) init(ab *AudioBuffer) error {
	ctx, err := oto.NewContext(
		int(ab.SamplesPerSecond),
		int(ab.ChannelCount),
		int(ab.BitsPerSample/8),
		int(ab.BlockSize))
	if err != nil {
		panic(err)
	}
	player := ctx.NewPlayer()
	ao.ctx = ctx
	ao.player = player
	ao.ab = ab
	return nil
}

func (ao *audioOutput) close() error {
	return ao.player.Close()
}

// handled in write
func (ao *audioOutput) noLongerBusy(blockIndex int) bool {
	return !ao.ab.blocks[blockIndex].busy
}

func (ao *audioOutput) write(bytes []byte, i int) {
	ao.player.Write(bytes)
	ao.ab.blocks[i].busy = false
}
