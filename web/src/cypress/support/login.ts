function normalizeURL(url: string | null): string {
  if (!url) throw new Error('url required')
  return new URL(url).toString()
}

function login(
  username?: string,
  password?: string,
  tokenOnly = false,
): Cypress.Chainable<string> {
  if (!username) {
    return cy
      .fixture('profile')
      .then((p) => login(p.username, p.password, tokenOnly))
  }
  if (!password) {
    return cy
      .fixture('profile')
      .then((p) => login(username, p.password, tokenOnly))
  }

  return cy
    .request({
      url: '/api/v2/identity/providers/basic?noRedirect=1',
      method: 'POST',
      form: true, // indicates the body should be form urlencoded and sets Content-Type: application/x-www-form-urlencoded headers
      body: {
        username,
        password,
      },
      followRedirect: false,
      headers: {
        referer: Cypress.config('baseUrl'),
        Cookie: '',
      },
    })
    .then((res) => {
      const token = res.body as string
      if (!tokenOnly) {
        cy.clearCookies()
        cy.setCookie('goalert_session.2', token)
        return ''
      }
      return token
    })
}

function adminLogin(tokenOnly = false): Cypress.Chainable<string> {
  return cy
    .fixture('profileAdmin')
    .then((p) => login(p.username, p.password, tokenOnly))
}

Cypress.Commands.add('login', login)
Cypress.Commands.add('adminLogin', adminLogin)

export {}
