import React from 'react'
import ReactDOM from 'react-dom'
import File from './File'

it('renders without crashing', () => {
  /* arrange */
  const div = document.createElement('div')

  /* act/assert */
  ReactDOM.render(
    <File
      name=''
      opRef=''
      param={{ description: '' }}
      value=''
    />,
    div)
})
