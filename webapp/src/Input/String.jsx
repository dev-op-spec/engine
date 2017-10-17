import React from 'react';
import Input from './Input';
import Select from './Select';
import opspecDataValidator from '@opspec/sdk/lib/data/string/validator';

export default ({name, onInvalid, onValid, string}) => {
  if (string.constraints && !string.isSecret && string.constraints.enum) {
    return <Select
      description={string.description}
      name={name}
      options={string.constraints.enum.map(item => ({name: item, value: item}))}
      onInvalid={onInvalid}
      onValid={value => onValid({string: value})}
      validate={value => opspecDataValidator.validate(value, string.constraints)}
      value={string.default}
    />
  }
  return <Input
    description={string.description}
    name={name}
    onInvalid={onInvalid}
    onValid={value => onValid({string: value})}
    type={string.isSecret ? 'password' : 'text'}
    validate={value => opspecDataValidator.validate(value, string.constraints)}
    value={string.default}
  />
};
