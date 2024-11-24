export ROOT_MOD=github.com/HCH1212/gomall

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service user -module ${ROOT_MOD}/rpc_gen -I ../idl && cd ../app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user -module ${ROOT_MOD}/app/user -I ../../idl --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY: docker
docker:
	@sudo docker compose up -d
