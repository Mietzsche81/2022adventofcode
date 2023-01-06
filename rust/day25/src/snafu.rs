use std::ops;

#[derive(Clone, Copy, Debug)]
pub struct Base5 {
    val: usize,
}
#[derive(Clone, Debug)]
pub struct Snafu {
    val: String,
}
#[derive(Clone, Copy, Debug)]
pub struct Base10 {
    val: usize,
}

pub trait Conversion {
    fn to_base10(&self) -> Base10;
    fn to_base5(&self) -> Base5;
    fn to_snafu(&self) -> Snafu;
    fn new(val: usize) -> Self;
}

impl Conversion for Base10 {
    fn to_base10(&self) -> Base10 {
        return *self;
    }

    fn to_base5(&self) -> Base5 {
        let mut out = Base5 { val: 0 };
        let mut dividend = self.val;
        let mut power: u32 = 0;
        while dividend > 0 {
            out.val += (dividend % 5) * 10usize.pow(power);
            dividend /= 5;
            power += 1;
        }
        return out;
    }

    fn to_snafu(&self) -> Snafu {
        return self.to_base5().to_snafu();
    }

    fn new(val: usize) -> Self {
        return Self { val: val };
    }
}

impl Conversion for Base5 {
    fn to_base10(&self) -> Base10 {
        let mut out = Base10 { val: 0 };
        let mut dividend = self.val;
        let mut power: u32 = 0;
        while dividend > 0 {
            out.val += (dividend % 10) * 5usize.pow(power);
            dividend /= 10;
            power += 1;
        }
        return out;
    }

    fn to_base5(&self) -> Base5 {
        return *self;
    }

    fn to_snafu(&self) -> Snafu {
        let mut out = Snafu {
            val: "".to_string(),
        };
        let mut dividend = self.val;
        let carry = |d: &mut usize| {
            // handle previous carry
            if (*d % 10) > 4 {
                *d -= 5;
                *d += 10;
            }
            let r = *d % 5;
            // If = or -, need to carry
            if r > 2 {
                *d += 10;
            }
        };
        while dividend > 0 {
            match dividend % 5 {
                0 => out.val = ["0", out.val.as_str()].join(""),
                1 => out.val = ["1", out.val.as_str()].join(""),
                2 => out.val = ["2", out.val.as_str()].join(""),
                3 => out.val = ["=", out.val.as_str()].join(""),
                4 => out.val = ["-", out.val.as_str()].join(""),
                _ => (), // literally impossible
            }
            carry(&mut dividend);
            dividend /= 10;
        }
        return out;
    }

    fn new(val: usize) -> Self {
        return Base10 { val: val }.to_base5();
    }
}

impl Conversion for Snafu {
    fn to_base10(&self) -> Base10 {
        let mut val: isize = 0;
        let mut i: u32 = 0;
        for c in self.val.chars().rev() {
            match c {
                '0' => val += 0 * 5isize.pow(i),
                '1' => val += 1 * 5isize.pow(i),
                '2' => val += 2 * 5isize.pow(i),
                '=' => val -= 2 * 5isize.pow(i),
                '-' => val -= 1 * 5isize.pow(i),
                _ => panic!("Invalid character '{}' in Snafu", c),
            }
            i += 1;
        }

        return Base10 { val: val as usize };
    }

    fn to_base5(&self) -> Base5 {
        return self.to_base10().to_base5();
    }

    fn to_snafu(&self) -> Snafu {
        return self.clone();
    }

    fn new(val: usize) -> Self {
        return Base10 { val: val }.to_snafu();
    }
}

impl ops::Add<&Snafu> for &Snafu {
    type Output = Snafu;

    fn add(self, rhs: &Snafu) -> Snafu {
        return Snafu::new(self.to_base10().val + rhs.to_base10().val);
    }
}

impl ops::Add<Snafu> for Snafu {
    type Output = Snafu;

    fn add(self, rhs: Snafu) -> Snafu {
        return Snafu::new(self.to_base10().val + rhs.to_base10().val);
    }
}

impl ops::Add<&Snafu> for Snafu {
    type Output = Snafu;

    fn add(self, rhs: &Snafu) -> Snafu {
        return Snafu::new(self.to_base10().val + rhs.to_base10().val);
    }
}

impl Snafu {
    pub fn parse(s: String) -> Snafu {
        return Snafu { val: s };
    }
}
