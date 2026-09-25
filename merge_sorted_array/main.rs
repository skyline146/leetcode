fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
    let (mut i, mut j) = (m as usize, n as usize); // i,j - counts how much elements left

    while j > 0 {
        let k = i + j - 1;
        if i > 0 && nums1[i - 1] >= nums2[j - 1] {
            nums1[k] = nums1[i - 1];
            i -= 1;
        } else {
            nums1[k] = nums2[j - 1];
            j -= 1;
        }
    }
}

fn main() {
    let mut nums1 = vec![1, 2, 3, 0, 0, 0];
    let mut nums2 = vec![2, 5, 6];
    merge(&mut nums1, 3, &mut nums2, 3);
}
