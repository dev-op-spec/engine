import React from 'react'
import jsYaml from 'js-yaml'
import TextArea from './AceEditor'
import ParamObject from '@opctl/sdk/src/types/param/object'
import paramObjectValidate from '@opctl/sdk/src/opspec/interpreter/opcall/params/param/object/validate'

interface Props {
  name: string
  object: ParamObject
  onInvalid?: () => any | null | undefined
  onValid: (value: any) => any
  opRef: string
  value: any
}

export default (
  {
    name,
    object,
    onInvalid,
    onValid,
    opRef,
    value
  }: Props
) => (
    <TextArea
      description={object.description}
      name={name}
      onInvalid={onInvalid}
      onValid={value => onValid(jsYaml.safeLoad(value))}
      opRef={opRef}
      validate={value => {
        try {
          return paramObjectValidate(jsYaml.safeLoad(value), object.constraints)
        } catch (err) {
          return [err]
        }
      }}
      value={value
        ? jsYaml.safeDump(value)
        : jsYaml.safeDump(object.default || '')
      }
    />
  )
