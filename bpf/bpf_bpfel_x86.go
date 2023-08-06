// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package bpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type BpfEvent struct {
	Pc   uint64
	Type uint8
	_    [7]byte
}

type BpfSkbmeta struct {
	Header [256]uint8
	Len    uint32
	Mark   uint32
}

// LoadBpf returns the embedded CollectionSpec for Bpf.
func LoadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Bpf: %w", err)
	}

	return spec, err
}

// LoadBpfObjects loads Bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*BpfObjects
//	*BpfPrograms
//	*BpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// BpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type BpfSpecs struct {
	BpfProgramSpecs
	BpfMapSpecs
}

// BpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type BpfProgramSpecs struct {
	BpfHelper *ebpf.ProgramSpec `ebpf:"bpf_helper"`
	OnEntry   *ebpf.ProgramSpec `ebpf:"on_entry"`
	OnExit    *ebpf.ProgramSpec `ebpf:"on_exit"`
}

// BpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type BpfMapSpecs struct {
	Events       *ebpf.MapSpec `ebpf:"events"`
	SkbLocations *ebpf.MapSpec `ebpf:"skb_locations"`
	Skbcaches    *ebpf.MapSpec `ebpf:"skbcaches"`
	Skbmetas     *ebpf.MapSpec `ebpf:"skbmetas"`
}

// BpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type BpfObjects struct {
	BpfPrograms
	BpfMaps
}

func (o *BpfObjects) Close() error {
	return _BpfClose(
		&o.BpfPrograms,
		&o.BpfMaps,
	)
}

// BpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type BpfMaps struct {
	Events       *ebpf.Map `ebpf:"events"`
	SkbLocations *ebpf.Map `ebpf:"skb_locations"`
	Skbcaches    *ebpf.Map `ebpf:"skbcaches"`
	Skbmetas     *ebpf.Map `ebpf:"skbmetas"`
}

func (m *BpfMaps) Close() error {
	return _BpfClose(
		m.Events,
		m.SkbLocations,
		m.Skbcaches,
		m.Skbmetas,
	)
}

// BpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type BpfPrograms struct {
	BpfHelper *ebpf.Program `ebpf:"bpf_helper"`
	OnEntry   *ebpf.Program `ebpf:"on_entry"`
	OnExit    *ebpf.Program `ebpf:"on_exit"`
}

func (p *BpfPrograms) Close() error {
	return _BpfClose(
		p.BpfHelper,
		p.OnEntry,
		p.OnExit,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_x86.o
var _BpfBytes []byte
