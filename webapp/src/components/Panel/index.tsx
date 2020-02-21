import React from 'react'
import { css } from 'emotion'
import brandColors from '../../brandColors'

/**
 * A panel
 */
export default props =>
  <div
    className={css({
      width: '100%',
      borderTop: `solid thin ${brandColors.reallyLightGray} !important`
    })}
  >
    {props.children}
  </div>
