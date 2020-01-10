---
title: Number
---

Number typed values are a number.

Numbers...
- are immutable, i.e. assigning to an number results in a copy of the original number
- can be passed in/out of ops via [number parameters](../structure/op-directory/op/parameter/number)
- can be initialized via [number initialization](#initialization)
- are coerced according to [number coercion](#coercion)

### Initialization
Number typed values can be constructed from a literal or templated number.
 
A templated number is a number which includes one or more value reference.
At runtime, each reference gets evaluated and replaced with it's corresponding value.

#### Initialization Example (literal)
```yaml
2
```

#### Initialization Example (templated)
given:
- `someNumber`
  - is in scope
   - is type coercible to number

```yaml
# $(someNumber) replaced w/ someNumber
222$(someNumber)3e10
```

### Coercion
Number typed values are coercible to:

- [boolean](#boolean) (numbers which are 0 coerce to false; all else coerce to true)
- [file](#file)
- [string](#string)

#### Coercion Example (number to file)
```yaml
name: numAsFile
run:
  container:
    image: { ref: alpine }
    cmd:
    - sh
    - -ce
    - cat /numCoercedToFile
    files:
      /numCoercedToFile: 2.2
```