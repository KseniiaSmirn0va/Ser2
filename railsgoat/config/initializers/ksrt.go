DetectorType_JiraToken                     DetectorType = 120
Token 23417846781247eef


			want: []detectors.Result{
				{
					DetectorType: detectorspb.DetectorType_AWS,
					Verified:     true,
					Redacted:     "AKIASP2TPHJSQH3FJRUX",
					ExtraData: map[string]string{
3 months ago

added resource type mapping to extraData in AWS (#2087)
						"resource_type":  "Access key",
						"rotation_guide": "https://howtorotate.com/docs/tutorials/aws/",
						"account":        "171436882533",
						"arn":            "arn:aws:iam::171436882533:user/canarytokens.com@@4dxkh0pdeop3bzu9zx5wob793",
						"user_id":        "AIDASP2TPHJSUFRSTTZX4",
5 months ago

Retry AWS verification 403s (#1757)
					},
				},
			},
