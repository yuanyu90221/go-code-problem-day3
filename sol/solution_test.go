package sol

import "testing"

func TestWidthOfBinaryTree(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				node: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 3,
							},
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WidthOfBinaryTree(tt.args.node); got != tt.want {
				t.Errorf("WidthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
