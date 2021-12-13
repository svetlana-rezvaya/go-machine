# Instruction Set

## "Load Constant" Instruction

Opcode: `0x1`.

Parameters:

- constant;
- register index.

Example:

![](load_instruction.svg)

## "Load Memory" Instruction

Opcode: `0x2`.

Parameters:

- memory index;
- register index.

Example:

![](load_instruction.svg)

## "Store Constant" Instruction

Opcode: `0x3`.

Parameters:

- constant;
- memory index.

Example:

![](store_instruction.svg)

## "Store Memory" Instruction

Opcode: `0x4`.

Parameters:

- register index;
- memory index.

Example:

![](store_instruction.svg)

## "Addition" Instruction

Opcode: `0x5`.

Parameters:

- left register index;
- right register index;
- result register index.

Example:

![](addition_instruction.svg)

## "Subtraction" Instruction

Opcode: `0x6`.

Parameters:

- left register index;
- right register index;
- result register index.

Example (in a similar way):

![](addition_instruction.svg)

## "Multiplication" Instruction

Opcode: `0x7`.

Parameters:

- left register index;
- right register index;
- result register index.

Example (in a similar way):

![](addition_instruction.svg)

## "Division" Instruction

Opcode: `0x8`.

Parameters:

- left register index;
- right register index;
- result register index.

Example (in a similar way):

![](addition_instruction.svg)

## "Modulo" Instruction

Opcode: `0x9`.

Parameters:

- left register index;
- right register index;
- result register index.

Example (in a similar way):

![](addition_instruction.svg)

## "Jump" Instruction

Opcode: `0xa`.

Parameters:

- IP register value.

Example:

![](jump_instruction.svg)

## "Jump If Negative" Instruction

Opcode: `0xb`.

Parameters:

- register index;
- IP register value.

Example (in a similar way):

![](jump_if_zero_instruction.svg)

## "Jump If Zero" Instruction

Opcode: `0xc`.

Parameters:

- register index;
- IP register value.

Example:

![](jump_if_zero_instruction.svg)

## "Jump If Positive" Instruction

Opcode: `0xd`.

Parameters:

- register index;
- IP register value.

Example (in a similar way):

![](jump_if_zero_instruction.svg)
