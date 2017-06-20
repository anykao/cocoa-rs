extern crate serde_json;
extern crate serde;
#[macro_use]
extern crate serde_derive;
#[macro_use]
extern crate log;
extern crate log4u;
extern crate linked_hash_map;

use std::fs;
use std::io::{self, Write, BufReader};
use std::io::prelude::*;
use serde_json::{Value, Error};
use std::collections::HashMap;

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
struct definition {
  aliases: HashMap<String, String>,
  called_definitions: Value,
  cftypes: Value, // empty object
  classes: HashMap<String, klass>,
  #[serde(rename = "enum")]
  enumeration: HashMap<String, Value>,
  expressions: Value, // empty object
  externs: HashMap<String, Value>,
  formal_protocols: HashMap<String, protocol>,
  func_macros: Value, // empty object
  functions: HashMap<String, function>,
  informal_protocols: HashMap<String, protocol>,
  literals: HashMap<String, Value>,
  structs: HashMap<String, strct>,
}

#[derive(Deserialize)]
struct strct {
  fieldnames: Vec<Value>,
  special: bool,
  typestr: String,
}

#[derive(Deserialize)]
struct protocol {
  implements: Vec<String>,
  methods: Vec<method>,
  properties: Value,
}
#[derive(Deserialize)]
struct function {
  args: Vec<arg>,
  retval: retval,
}

#[derive(Deserialize)]
struct arg {
  typestr: String,
  typestr_special: Option<bool>,
}

#[derive(Deserialize)]
struct retval {
  typestr: String,
  typestr_special: Option<bool>,
}

#[derive(Deserialize)]
struct method {
  args: Vec<arg>,
  class_method: bool,
  retval: retval,
  selector: String,
  visibility: String,
}

#[derive(Deserialize)]
struct klass {
  methods: Vec<method>,
  name:       String,
  properties: Value,
  protocols:  Value,
}

#[derive(Deserialize)]
struct fwinfo {
  arch: String,
  definitions: definition,
  framework: String,
  headers: Vec<String>,
  release: String,
  sdk: String,
}

fn main() {
  log4u::init().unwrap();

  let f = fs::File::open("x86_64-10.10.fwinfo").unwrap();
  let f = BufReader::new(f);
  let mut lines = f.lines().skip(4);

  let body: Vec<String> = lines.map(|line|line.unwrap()).collect();


  let v: Value = serde_json::from_str(&body.join("\n")).unwrap();

  process(v);

}

fn inspect(v: Value) {
  match v {
    Value::Null => debug!("null"),
    Value::Bool(b) => debug!("Bool {}",b),
    Value::Number(n) => debug!("Number {}",n),
    Value::String(s) => debug!("String {}", s),
    Value::Array(a) => debug!("Array {:?}", a),
    Value::Object(m) => {
      for (k,v) in m {
        debug!("Object {}", k)
      }
    },
  }
}


fn process(v: Value) {
  let info: fwinfo = serde_json::from_value(v).unwrap();
  for (k, _ ) in info.definitions.classes {
    info!("{}",k);
  }
}
