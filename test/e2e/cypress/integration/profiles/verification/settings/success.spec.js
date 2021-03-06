import { APP_URL, assertVerifiableAddress, gen } from '../../../../helpers'

context('Settings', () => {
  describe('successful flow', () => {
    let identity

    before(() => {
      cy.deleteMail()
    })

    beforeEach(() => {
      identity = gen.identity()
      cy.register(identity)
      cy.deleteMail({ atLeast: 1 }) // clean up registration email

      cy.login(identity)
      cy.visit(APP_URL + '/settings')
    })

    it('should update the verify address and request a verification email', () => {
      const email = `not-${identity.email}`
      cy.get('#user-profile input[name="traits.email"]').clear().type(email)
      cy.get('#user-profile button[type="submit"]').click()
      cy.get('.container').should(
        'contain.text',
        'Your changes have been saved!'
      )
      cy.get('#user-profile input[name="traits.email"]').should(
        'contain.value',
        email
      )
      cy.session().then(assertVerifiableAddress({ isVerified: false, email }))

      cy.verifyEmail({ expect: { email } })

      cy.location('pathname').should('eq', '/')
    })

    xit('should should be able to allow or deny (and revert?) the address change', () => {
      // FIXME https://github.com/zzpu/ums/issues/292
    })
  })
})
