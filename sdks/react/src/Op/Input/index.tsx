import React from 'react'
import ModelParam from '@opctl/sdk/src/model/param'
import InputArray from './Array'
import InputBoolean from './Boolean'
import InputDir from './Dir'
import InputFile from './File'
import InputNumber from './Number'
import InputObject from './Object'
import InputSocket from './Socket'
import InputString from './String'
import Value from '@opctl/sdk/src/model/value'

interface DataStore {
  get(
    dataRef: string
  ): Promise<Value | null | undefined>
}

interface Props {
  dataStore?: DataStore | null | undefined
  input: ModelParam
  name: string
  onInvalid?: () => any | null | undefined
  onValid: (value: any) => any
  opRef: string
  value: any
}

export default (
  {
    dataStore,
    input,
    name,
    onInvalid,
    onValid,
    opRef,
    value
  }: Props
) => {
  if (dataStore) {
    value = value || dataStore.get(name)
  }

  // delegate to component for input
  if (input.array) {
    return <InputArray
      array={input.array}
      name={name}
      onInvalid={onInvalid}
      onValid={onValid}
      opRef={opRef}
      value={value}
    />
  } else if (input.boolean) {
    return <InputBoolean
      boolean={input.boolean}
      name={name}
      onValid={onValid}
      opRef={opRef}
      value={value}
    />
  } else if (input.dir) {
    return <InputDir
      dir={input.dir}
      name={name}
      onValid={onValid}
      opRef={opRef}
      value={value}
    />
  } else if (input.file) {
    return <InputFile
      file={input.file}
      name={name}
      onValid={onValid}
      opRef={opRef}
      value={value}
    />
  } else if (input.number) {
    return <InputNumber
      onInvalid={onInvalid}
      onValid={onValid}
      name={name}
      number={input.number}
      opRef={opRef}
      value={value}
    />
  } else if (input.object) {
    return <InputObject
      name={name}
      object={input.object}
      onInvalid={onInvalid}
      onValid={onValid}
      opRef={opRef}
      value={value}
    />
  } else if (input.socket) {
    return <InputSocket
      name={name}
      onValid={onValid}
      socket={input.socket}
      opRef={opRef}
      value={value}
    />
  } else if (input.string) {
    return <InputString
      name={name}
      onInvalid={onInvalid}
      onValid={onValid}
      string={input.string}
      opRef={opRef}
      value={value}
    />
  }
  return null
}
