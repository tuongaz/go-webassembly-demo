extern crate rand;

use rand::Rng;
use serde_json::json;
use std::ffi::c_int;

pub struct Food {
    pub image: char,
    pub coordination: (i32, i32),
}

impl Food {
    pub fn new(width: i32, height: i32) -> Self {
        let mut rng = rand::thread_rng();

        let emojis = vec!['🍏', '🍌', '🍑', '🍊', '🍇', '🥝', '🥑', '🍓'];

        let x: i32 = rng.gen_range(4..width - 4);
        let y: i32 = rng.gen_range(4..height - 4);

        Food {
            image: emojis[rng.gen_range(0..emojis.len())],
            coordination: (x, y),
        }
    }
}

#[no_mangle]
pub extern "C" fn create_food(width: i32, height: i32) -> u64 {
    let food = Food::new(width, height);
    let food_json = json!({
        "image": food.image ,
        "coordination": {
            "x": food.coordination.0,
            "y": food.coordination.1,
        },
    });

    // show diagram
    let string = food_json.to_string();
    let ptr = string.as_ptr() as u64;
    let len = string.len() as c_int;

    std::mem::forget(string);

    return (ptr << 32) | len as u64;
}

