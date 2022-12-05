// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::*;
use std::{any::Any, sync::mpsc, thread::spawn};
use wasmlib::*;

// TODO to handle the request in parallel, WasmClientContext must be static now.
// We need to solve this problem. By copying the vector of event_handlers, we may solve this problem
pub struct WasmClientContext {
    pub chain_id: ScChainID,
    pub event_handlers: Vec<Box<dyn IEventHandlers>>,
    pub key_pair: Option<keypair::KeyPair>,
    pub req_id: ScRequestID,
    pub sc_name: String,
    pub sc_hname: ScHname,
    pub svc_client: WasmClientService, //TODO Maybe  use 'dyn IClientService' for 'svc_client' instead of a struct
}

impl WasmClientContext {
    pub fn new(
        svc_client: &WasmClientService,
        chain_id: &wasmlib::ScChainID,
        sc_name: &str,
    ) -> WasmClientContext {
        WasmClientContext {
            svc_client: svc_client.clone(),
            sc_name: sc_name.to_string(),
            sc_hname: ScHname::new(sc_name),
            chain_id: chain_id.clone(),
            event_handlers: Vec::new(),
            key_pair: None,
            req_id: request_id_from_bytes(&[]),
        }
    }

    pub fn default() -> WasmClientContext {
        WasmClientContext {
            svc_client: WasmClientService::default(),
            sc_name: String::new(),
            sc_hname: ScHname(0),
            chain_id: chain_id_from_bytes(&[]),
            event_handlers: Vec::new(),
            key_pair: None,
            req_id: request_id_from_bytes(&[]),
        }
    }

    pub fn current_chain_id(&self) -> ScChainID {
        return self.chain_id;
    }

    pub fn init_func_call_context(&'static self) {
        wasmlib::host::connect_host(self);
    }

    pub fn init_view_call_context(&'static self, _contract_hname: &ScHname) -> ScHname {
        wasmlib::host::connect_host(self);
        return self.sc_hname;
    }

    pub fn register(&'static mut self, handler: Box<dyn IEventHandlers>) -> errors::Result<()> {
        for h in self.event_handlers.iter() {
            if handler.type_id() == h.as_ref().type_id() {
                return Ok(());
            }
        }
        self.event_handlers.push(handler);
        if self.event_handlers.len() > 1 {
            return Ok(());
        }
        return self.start_event_handlers();
    }

    // overrides default contract name
    pub fn service_contract_name(&mut self, contract_name: &str) {
        self.sc_hname = wasmlib::ScHname::new(contract_name);
    }

    pub fn sign_requests(&mut self, key_pair: &keypair::KeyPair) {
        self.key_pair = Some(key_pair.clone());
    }

    pub fn unregister(&mut self, handler: Box<dyn IEventHandlers>) {
        self.event_handlers.retain(|h| {
            if handler.type_id() == h.as_ref().type_id() {
                return false;
            } else {
                return true;
            }
        });
        if self.event_handlers.len() == 0 {
            self.stop_event_handlers();
        }
    }

    pub fn wait_request(&mut self, req_id: Option<&ScRequestID>) -> errors::Result<()> {
        let r_id;
        match req_id {
            Some(id) => r_id = id,
            None => r_id = &self.req_id,
        }
        return self.svc_client.wait_until_request_processed(
            &self.chain_id,
            &r_id,
            std::time::Duration::new(60, 0),
        );
    }

    pub fn start_event_handlers(&'static self) -> errors::Result<()> {
        let (tx, rx): (mpsc::Sender<String>, mpsc::Receiver<String>) = mpsc::channel();
        self.svc_client.subscribe_events(tx).unwrap();

        spawn(move || {
            for msg in rx {
                self.process_event(&msg).unwrap();
            }
        });

        return Ok(());
    }

    pub fn stop_event_handlers(&self) {
        todo!()
    }

    fn process_event(&self, _msg: &str) -> errors::Result<()> {
        // FIXME parse the msg
        for handler in self.event_handlers.iter() {
            handler
                .as_ref()
                .call_handler("topic", &vec!["params".to_string()]); // FIXME use the correct parsed message
        }
        todo!()
    }
}

#[cfg(test)]
mod tests {
    use crate::*;
    use wasmlib::*;

    struct FakeEventHandler {}
    impl IEventHandlers for FakeEventHandler {
        fn call_handler(&self, _topic: &str, _params: &Vec<String>) {}
    }

    #[test]
    fn test_wasm_client_context_new() {
        let svc_client = wasmclientservice::WasmClientService::default();
        let chain_id = wasmlib::chain_id_from_bytes(&vec![
            41, 180, 220, 182, 186, 38, 166, 60, 91, 105, 181, 183, 219, 243, 200, 162, 131, 181,
            57, 142, 41, 30, 236, 92, 178, 1, 116, 229, 174, 86, 156, 210,
        ]);
        let sc_name = "sc_name";
        let ctx = wasmclientcontext::WasmClientContext::new(&svc_client, &chain_id, sc_name);
        assert!(ctx.svc_client == svc_client);
        assert!(ctx.sc_name == sc_name);
        assert!(ctx.sc_hname == wasmlib::ScHname::new(sc_name));
        assert!(ctx.chain_id == chain_id);
        assert!(ctx.event_handlers.len() == 0);
        // assert!(ctx.key_pair == None);
        assert!(ctx.req_id == wasmlib::request_id_from_bytes(&[]));
    }

    #[test]
    fn test_register() {}

    #[test]
    fn test_call_view_by_hname() {}
}