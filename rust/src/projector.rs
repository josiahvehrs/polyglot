use crate::config::Config;
use serde::{Deserialize, Serialize};
use std::{collections::HashMap, path::PathBuf, vec};

#[derive(Debug, Serialize, Deserialize)]
struct Data {
    pub projector: HashMap<PathBuf, HashMap<String, String>>,
}

pub struct Projector {
    config: Config,
    data: Data,
}

fn default_data() -> Data {
    Data {
        projector: HashMap::new(),
    }
}

impl Projector {
    pub fn get_value_all(&self) -> HashMap<&String, &String> {
        let curr = Some(self.config.pwd.as_path());
        let mut paths = vec![];
        let mut out = HashMap::new();

        while let Some(p) = curr {
            paths.push(p);
        }

        for path in paths.into_iter().rev() {
            if let Some(x) = self.data.projector.get(path) {
                out.extend(x.iter())
            }
        }

        out
    }

    pub fn get_value(&self, key: &str) -> Option<&String> {
        let mut curr = Some(self.config.pwd.as_path());
        let mut out = None;

        while let Some(p) = curr {
            if let Some(dir) = self.data.projector.get(p) {
                if let Some(v) = dir.get(key) {
                    out = Some(v);
                    break;
                }
            }
            curr = p.parent();
        }

        out
    }

    pub fn set_value(&mut self, key: String, value: String) {
        self.data
            .projector
            .entry(self.config.pwd.clone())
            .or_default()
            .insert(key, value);
    }

    pub fn remove_value(&mut self, key: &str) {
        self.data
            .projector
            .entry(self.config.pwd.clone())
            .or_default()
            .remove(key);
    }

    pub fn from_config(config: Config) -> Self {
        if std::fs::metadata(&config.config).is_ok() {
            let contents = std::fs::read_to_string(&config.config);
            let contents = contents.unwrap_or("{\"projector\": {}}".to_string());
            let data = serde_json::from_str(&contents).unwrap_or(default_data());

            return Projector { config, data };
        }
        Projector {
            config,
            data: default_data(),
        }
    }
}

#[cfg(test)]
mod test {
    use collection_macros::hashmap;
    use std::{collections::HashMap, path::PathBuf};

    use crate::config::Config;

    use super::{Data, Projector};

    fn get_data() -> HashMap<PathBuf, HashMap<String, String>> {
        hashmap! {
                PathBuf::from("/") => hashmap! {
                    "foo".into() => "bar1".into(),
                    "fem".into() => "is_great".into(),
                },
                PathBuf::from("/foo") => hashmap!{
                    "foo".into() => "bar2".into(),
                },
                PathBuf::from("/foo/bar") => hashmap!{
                    "foo".into() => "bar3".into(),
                },
        }
    }

    fn get_projector(pwd: PathBuf) -> Projector {
        Projector {
            config: Config {
                config: PathBuf::from(""),
                pwd,
                operation: crate::config::Operation::Print(None),
            },
            data: Data {
                projector: get_data(),
            },
        }
    }

    #[test]
    fn get_value() {
        let proj = get_projector(PathBuf::from("/foo/bar"));
        assert_eq!(proj.get_value("foo"), Some(&String::from("bar3")));
        assert_eq!(proj.get_value("fem"), Some(&String::from("is_great")));
    }

    #[test]
    fn set_value() {
        let mut proj = get_projector(PathBuf::from("/foo/bar"));
        proj.set_value("foo".into(), "bar4".into());
        proj.set_value("fem".into(), "is_better_than_great".into());
        assert_eq!(proj.get_value("foo"), Some(&String::from("bar4")));
        assert_eq!(
            proj.get_value("fem"),
            Some(&String::from("is_better_than_great"))
        );
    }

    #[test]
    fn remove_value() {
        let mut proj = get_projector(PathBuf::from("/foo/bar"));
        proj.remove_value("foo");
        proj.remove_value("fem");

        assert_eq!(proj.get_value("foo"), Some(&String::from("bar2")));
        assert_eq!(proj.get_value("fem"), Some(&String::from("is_great")));
    }
}
