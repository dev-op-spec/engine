import React from 'react'
import Markdown from '../Markdown'
import 'highlightjs/styles/github.css'

export default ({ value, opRef }) =>
  <div className='custom-control-description'>
    <Markdown value={value} opRef={opRef} />
  </div>
