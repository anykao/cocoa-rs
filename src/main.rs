extern crate serde_json;
extern crate serde;
#[macro_use]
extern crate serde_derive;
#[macro_use]
extern crate log;
extern crate log4u;

use std::fs;
use std::io::{self, Write, BufReader};
use std::io::prelude::*;
use serde_json::{Value, Error};


#[derive(Deserialize)]
struct definition {
  aliases: String,
  called_definitions: String,
  cftypes: String,
  classes: String,
  #[serde(rename = "enum")]
  enumeration: String,
  expressions: String,
  externs: String,
  formal_protocols: String,
  func_macros: String,
  informal_protocols: String,
  literals: String,
  structs: String,
}


struct arg {
  typestr: String,
  typestr_special: String,
}
struct retval {
  typestr: String,
  typestr_special: String,
}
struct method {
  args: Vec<arg>,
  class_method: bool,
  retval: retval,
  selector: String,
  visibility: String,
}

struct klass {
  methods: Vec<method>,
  name:       String,
  properties: String,
  protocols:  String,
}


#[derive(Deserialize)]
struct fwinfo {
  arch: String,
  definitions: Value,
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

  println!("{}",&body.join("\n"));

//  let v: Value = serde_json::from_str(&body.join("\n")).unwrap();


//  inspect(v);


//  process(v);

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

  inspect(info.definitions);

}
