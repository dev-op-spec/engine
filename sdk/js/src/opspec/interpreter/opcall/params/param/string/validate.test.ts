import objectUnderTest from './validate'

describe('validate', () => {
  describe('constraints undefined', () => {
    it('returns expected result', () => {
      /* arrange/act */
      const actualErrors = objectUnderTest(
        'dummyValue'
      )

      /* assert */
      expect(actualErrors).toEqual([])
    })
  })
  describe('constraints defined', () => {
    describe('format constraint', () => {
      describe('constraint.format is docker-image-ref', () => {
        describe('value is docker-image-ref', () => {
          it('returns expected result', () => {
            /* arrange/act */
            const actualErrors = objectUnderTest('owner/registry:tag', {
              format: 'docker-image-ref'
            })

            /* assert */
            expect(actualErrors).toEqual([])
          })
        })
        describe('value isnt docker-image-ref', () => {
          it('returns expected result', () => {
            /* arrange/act */
            const actualErrors = objectUnderTest('#', {
              format: 'docker-image-ref'
            })

            /* assert */
            expect(actualErrors).toEqual([{
              'dataPath': '',
              'keyword': 'format',
              'message': 'should match format "docker-image-ref"',
              'params': {'format': 'docker-image-ref'},
              'schemaPath': '#/format'
            }])
          })
        })
      })
    })
    describe('minLength constraint', () => {
      describe('value > minLength', () => {
        it('returns expected result', () => {
          /* arrange/act */
          const actualErrors = objectUnderTest(
            'dummyValue',
            {
              minLength: 1
            }
          )

          /* assert */
          expect(actualErrors).toEqual([])
        })
      })
      describe('value < minLength', () => {
        it('returns expected result', () => {
          /* arrange/act */
          const actualErrors = objectUnderTest(
            'dummyValue',
            {
              minLength: 100
            }
          )

          /* assert */
          expect(actualErrors).toEqual([{
            'dataPath': '',
            'keyword': 'minLength',
            'message': 'should NOT be shorter than 100 characters',
            'params': {'limit': 100},
            'schemaPath': '#/minLength'
          }])
        })
      })
    })
  })
})
