// cross-env NODE_ENV=testing knex migrate:latest ; cross-env NODE_ENV=testing mocha --timeout=30000 --exit
const chai = require('chai')
const chaiHttp = require('chai-http')
const server = require('../src/main')

chai.should()
chai.use(chaiHttp)

describe("Users test suite", _ => {

  let req = chai.request(server.app)

  let user = {
    loginuser: "mrbigode",
    hashuser: "f447b20a7fcbf53a5d5be013ea0b15af"//echo 123456 | md5sum
  }

  let iduser = 0

  it("should create one user", done => {
    req.post("/user/save").send(user).end((err, ret) => {
      ret.should.have.status(200)
      iduser = ret.body.iduser
      done(err)
    })
  })

  it("should list at least one user", done => {
    req.get("/user/list").end((err, ret) => {
      ret.body.should.have.lengthOf.above(0)
      done(err)
    })
  })

  it("should update user", done => {
    const loginuser = "msverao"
    req.put("/user/save").send({ iduser, loginuser }).end((err, ret) => {
      ret.body.loginuser.should.be.equal(loginuser)
      done(err)
    })
  })

  const idrole = 2

  it("should add customer role to the user", done => {
    req.post("/user/addrole").send({ iduser, idrole }).end((err, ret) => {
      req.get(`/user/${iduser}`).end((err, ret) => {
        ret.body.roles.should.have.lengthOf.above(0)
        done(err)
      })
    })
  })

  it("should delete customer role to the user", done => {
    req.post("/user/removerole").send({ iduser, idrole }).end((err, ret) => {
      req.get(`/user/${iduser}`).end((err, ret) => {
        ret.body.roles.should.have.lengthOf(0)
        done(err)
      })
    })
  })

  it("should delete one user", done => {
    req.del(`/user/${iduser}`).end((err, ret) => {
      ret.status.should.be.equal(200)
      done(err)
    })
  })
})