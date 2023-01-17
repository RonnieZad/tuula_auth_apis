# Authentication Microservoce

### Project Backlog
User should be able to register with email address, phone_number, password, repeat_password and roles
User contact details are verified with sms OTP and email verification ID or token
User is required to setup security question after account registration
User account gets locked when user tries logging in more than 5 times
User can reset account password using the security question and security answer
User account should be locked for 2 hours after 6 times failed logins
Staff account can not be registered through guest registration
Only Staff account registration happens internally through the web dashboard
Account can not be deleted, but can be deactivated


### END POINTS TO DEVELOP
- [x] account-registration
- [x] account-login
- [ ] verify-account-with-otp
- [ ] verify-account-with-id
- [ ] reset-account-password
- [ ] update-account
- [ ] deactivate-account
- [ ] activate-account
- [ ] lock-account
- [ ] unlock-account
- [ ] logout
- [ ] register-new-landlord-account
- [ ] register-new-agent-account