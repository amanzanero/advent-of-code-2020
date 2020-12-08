package vmachine

import "errors"

type VMachine struct {
	acc, instruction int
	Ops              []Op
	history          map[int]bool
	corrupted        bool
}

func CreateVMachine() *VMachine {
	vm := VMachine{acc: 0, instruction: 0, corrupted: false}
	vm.Ops = make([]Op, 0)
	vm.history = make(map[int]bool)
	return &vm
}

func (vm *VMachine) Copy() *VMachine {
	cpyOps := make([]Op, len(vm.Ops))
	cpyHistory := make(map[int]bool)

	copy(cpyOps, vm.Ops)
	for k, v := range vm.history {
		cpyHistory[k] = v
	}

	return &VMachine{
		acc:         vm.acc,
		instruction: vm.instruction,
		Ops:         cpyOps,
		history:     cpyHistory,
		corrupted:   vm.corrupted,
	}
}

func (vm *VMachine) AddOp(op Op) {
	vm.Ops = append(vm.Ops, op)
}

func (vm *VMachine) NextOp() Op {
	return vm.Ops[vm.instruction]
}

func (vm *VMachine) IsCorrupted() bool {
	return vm.corrupted
}

func (vm *VMachine) RunProgram() {
	for true {
		if vm.instruction == len(vm.Ops) {
			break
		} else if vm.instruction > len(vm.Ops) || vm.instruction < 0 {
			vm.corrupted = true
			break
		}

		op := vm.NextOp()
		err := op.Execute(vm)

		if err != nil {
			vm.corrupted = true
			break
		}
	}
}

func (vm *VMachine) Reset() {
	vm.acc = 0
	vm.instruction = 0
	vm.history = make(map[int]bool)
	vm.corrupted = false
}

func (vm *VMachine) AddHistory() error {
	if vm.history[vm.instruction] {
		return errors.New("loop detected")
	}
	vm.history[vm.instruction] = true
	return nil
}

func (vm *VMachine) Acc() int {
	return vm.acc
}

func (vm *VMachine) ReplaceJmpAt(n int) {
	prevOp := vm.Ops[n]
	vm.Ops[n] = &JmpOp{opType: "jmp", value: prevOp.Value()}
}

func (vm *VMachine) ReplaceNOpAt(n int) {
	prevOp := vm.Ops[n]
	vm.Ops[n] = &NOp{opType: "nop", value: prevOp.Value()}
}
