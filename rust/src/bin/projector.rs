use clap::Parser;
use polyglot::opts;

fn main() {
    let opts = opts::Opts::parse();
    println!("{:?}", opts);
}
