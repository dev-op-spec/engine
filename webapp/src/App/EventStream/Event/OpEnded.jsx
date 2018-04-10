import React from 'react'

export default ({
  opEnded,
  timestamp
}) => {
  let color
  switch (opEnded.outcome) {
    case 'FAILED':
      color = 'rgb(255, 110, 103)'
      break
    case 'SUCCEEDED':
      color = 'rgb(95, 250, 104)'
      break
    case 'KILLED':
      color = 'rgb(96, 253, 255)'
      break
    default:
      throw new Error(`received unexpected OpEnded.Outcome: '${opEnded.outcome}'`)
  }

  return (
    <div style={{color}}>
      OpEnded
      Id='{opEnded.opId}'
      OpRef='{opEnded.opRef}'
      Outcome='{opEnded.outcome}'
      Timestamp='{timestamp}'
    </div>
  )
}
