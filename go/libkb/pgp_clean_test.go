package libkb

import (
	"testing"
	"encoding/base64"
)

func TestKeithWinstein(t *testing.T) {
	// Encode the key with spaces as a Base64-encoded string, since trailing
	// spaces don't survive our various text editors
	key64 := "LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tClZlcnNpb246IEdudVBHIHYxCgptUUlOQkU4dXorWUJFQURSbFZXSEZvem5TcC9vb010OGdIZDVZVGt5aGFQUVdCbXBaRzZFUkp0dGdMakF2TWRRCmE0ZlVVOTk0ZE5mdTBBVGNhbWVkUmNNdEtjQ0NWTTZETE1NWjNJVnhhWXg0SjY1OEJnVEcvK21iVnI5UkJRUUMKcXM4OWNRQ0dianBFRnZKSGNnc2lLNldNcVlWVXFvOUpxYzlEWFZlYThpVGpjQVV2a202OGFkMnBuRi9ZNUVFSwp3Rndaem43S3lkMEtoYmhlNmZCcElxeWVvcWJtL2lyMXpSd2JRWktJRzhqTGtER3hQMHlTcFh4VXFPVHBXYWxsCmQxcTF1SzdhT0pnTDYzWUZFdSs1VU5DU1Y1T2VabCtXb21yay96M0NwdmVyZ1BndHd3eEU4WlFOeU41MzR3VUgKZ1BLUkpRTEtGQjU4RE53WXJNRnZtdlVLT3J1UHU1czd4UDBCb2luOGRWRWcrRCtVMEN6UlNIcDU3UXQrNlNtOQpZOWpHcmRyQW1uZURhRjhuaWhlc1ZDUDNldG1IZEUxcXBZa0E1TlhKK3k4V2YyNWdxUjBRK21EbnE2QjlON1VJCkdCQk9jbDByWUZYSCtHNE52eG5OZUJjN3IxeHdiZkdKWUt6V21CV2ZwbzQ4bThWeXpKQ1lPSDBnMWJSRk4ydTYKU09JZ2JjbWRnNFFHVWhudHRDUllsYS9YYUV2QVA5cDlFWStLQWJCY2JtTEJZTXpXWi9wWm9KQ25xTUN1cWJvWQpkZTRmQXNyYWphdjFWOVNMVVN4VXNWQW4weUI1Q3lUaDNleWlqRUhxVUs3cC9mMk9wdGZObDU0dWZNTjd0eU0xCmZBaUNnTDdLL2JudVNDUWpDeGw4L1BORjVuR0Zhb0tGN2NpcTdxcVd6UUp5R2VmME9SSys4YURwZndBUkFRQUIKdEI5TFpXbDBhQ0JYYVc1emRHVnBiaUE4YTJWcGRHaDNRRzFwZEM1bFpIVStpUUkrQkJNQkFnQW9CUUpQTHMvbQpBaHNEQlFrRm81cUFCZ3NKQ0FjREFnWVZDQUlKQ2dzRUZnSURBUUllQVFJWGdBQUtDUkFndHlnNi9pVk1hZWhECkQvNDlUdG1LdmllNU5NaVAvZ1dOd3pDa3RmcXlZc3pBcmdWV0UraGNIWEQ4bWNUWHNwd2dGak5VbUlUUDRCNnEKaENkcnpsMXU5ZGhDVE5yYUJ5OUJQWFlIaWQwRzQ0Z1MxYmFjZVBGdWN2S3hMbHVRdzh5WGxhZVVacUdXd0tzKwpZZ2VqYzBIYjhTbjVjaDNtSFp3Nzkyc0RxWWxNRmRja1FHUlpsYVUxSUpzZVFqODYyaCtQWFNURmtHQ3I1bjA2CjlzMGQxU09IUUhyZnlQZGhCRVdPbWxTMVJNVGlXakZZTG5HcjRUUFc0OVJ4WFcraG5XL0xZbWIxbEdCQk9vT3MKdGJ6S3lHbjRFNGMrQXJTd2U5dW1KdWRLSDdyektseFg2UmNKSFZGUlRDUXladzQ1OEZ2ajA5NFB6WlhCL3Fwcwp1cXlhYzFYdHhCQnlqeHJ3M3hHM2RTQ3Q2bitZczZjaWlkeG9XcEV6SVFCazAwZnE0ak5nWThNbjhycmIvb1h0CjNVeTZVbktpdWpXbEg0UXFEd2RVdGFDUWNXQ010TFZ2ZVMzaksxbGg0T2lsZ0s3V0JXWnlOV0RwVFU3aVlabXgKeEU3cHlBZmV5cElNUkF5c29KWGNDa21MM1lkcjQybEc3TnpOeEZyNmpGMG1xYkozMTFEUU5vcVBKMkNPd3oySQpjVEsyYnpHNy90VWMwNVF1ZElTaGs0R2VsWTJLVVZ3Nzk2UVp4QnBISjZmOGxESkt0RldKd2NsQ2tDTFZBdisvCnkyS0JEb2JBWFJkUkZKcURYSTJGNHZsYWEyZHdZMURwbDA4Tjl5SlkzYVhici8xNHdOdHMwamJ6OG4wMEI3TEgKbkVNQ1ZxYXRCS1lpM1B1d2o3cHpHRlNTaEZiUVhEZkl4REQzVUg2aHhnbU5VSWtDSEFRUUFRZ0FCZ1VDVDhUcwpnd0FLQ1JCa2xaL3BnNDN4bkZIMUQvNDVBUkZXTHdwS2tRRHErdEVVSDJVT2hnZktSQ0VaVWNaYTYzdnI1T1ZtCldTVlhsLzZtcjMwbmNoS2FaUEdlbmE4aDRNWndiOGN1RGhXUVE5Q1ZRYURrVXV0cVVtUkVkajlPanNmN21WalMKTE9IRVBWS3U4TklZVThXSmw0NXpBT2I1ajR6QUpZTE94ZWhWeFQ1KzJIdnlLNVhZSndBTmlPV1pFTzVneERwWAozTC9hSVpOTHJSNWdZc3pZOVI3aU0zSHpFUHIvTkxJaktOYU9zSFB3VS8rZ2pKVDU4VXYzajh5WlRmWGptYWYvCnVOL25qOG4rV2pNMjVBdWM1UFhoZWJxWkh2M3ZSYUl0MWJEVWQ3dUZZei9lWFJTNVNkN0lDMDJjd0lxSURXZ2UKK01semo0NW5qdm5vemZ2Q3lTZ01ZWTVyODE4eFY2NVhCb0pXWVBBTGtpMDhyWGRCd3lobjZEbUhZYkptSW0vbQpHQVl6SDY5dHRodjZDNE9XYS8rZS9LNkF0d2o4aHExVlU1TXJpN3hOM2RURnVaK1greGxWL2pXMlZYYmVqbmQwClZ2aytUcHlzd2ZzZUpjNmc1MzVOMG1MakZ3V1hheE1GSDB2SUt6M0o1MWFuQlI5alo3d3RYeGQvTXRVcG9wZ0sKdXp1WmJxSnZWTk9CNkJydDBUTUZJaWdoK2tNZ2pLS2ZJVENNSUlWanFSQkV4U1JCYXRSOUw5TXZGY2Q3VHlkbQpPbWt2T0lYTXZnayswb0Rlb3AzRTVSelNQR1kyQUVlZnlXZXdaNUdoaHRFOEFoSC80WTZiVHhiZHRnc21laG1tCk5uTzhIZlBtZXROTlFGT0dMblQrbnBVRWZxcENtNjAvcmFyWmxZSmhKeFE5UzRsRHd6UDBrN3dBNEhibzAwM04KVjRoR0JCTVJDQUFHQlFKUHhZZnZBQW9KRUxIS2t1aW4yR3VWeDU4QW5pdyt6V3NWL3ZkMEhsOXorT2g4YnB6eQpSa20xQUtDT3pWRnVqREs3UUF6YnlSelZydFIwdytILzVJa0NIQVFUQVFnQUJnVUNUOFdKa1FBS0NSRG1OeGIwCjRHbVVubG42RUFEYXZOM2lJc3ZVSlBPZDB5MjRDaGMwemRwWWNyUTRZT0Rpa3ZFSHBiTjAwd0tFNjQ3MnhxSjEKV21XNW1uNlZTelFCQmg2RmtDZGN0eURnZVhJcE1LbzFzYSt4eitJR05XTGZvZUxGdmE2b1Q5QU01UDVmMGF2Ugo5QnFuSmJoS1lCaVowaG5qVUwrekVxZkEzYUxKQXVFQmdaS0VLK0lOV3BmMnk4MWpHZ2ZDR3hWOUVuREkzRCtoCmI5dnI5djZQMmRIMGx2ZGFPaEZseDVYSzhhVlc2VjgxTE1iK1lTQ2xJeEhwQVMrWjdla2xiS1dBNHpFZmI5N2IKOEJxdWJwQmpSQkJudzh6QStKd3FjS3pnd0pHckRQeGRKUHVpYUxaRWVvSzBZbS9seFkvOCtnMklub0JoMndLdwpDbFA4UHdyKzNnRVExN2FRNk1HMXI0VlF4YXI4TTlnL2w2MnlaaWVDbHdjbWpla0ZBM0w1U1NhVWlEbzdJakRjCnk1TzY1cm1UbVJ2Z0tmbkZ1amMrRStsUndaT0JURFI3dHgraHp1N25vTHhPaXpMYklsSWoweGZEWUZhRVBadmwKUzZjZjVuRmFRSnA3YVE2bkcrWGxVZkMvcVBwYzh0SGQ5dytXSHVyeU1sdE9sa0c0VzIreDE0elY4NWdSMkZXUQo2VEJETkRDUkhvcWExUGorVUhqOXpTNkp6Q0RmdVo2WUphS1IxNFhoZm5wVjVMNVlnd1N5VGQwL2Fka2QyaVdxClRaS2hOZXJVcGcxR3A2aUt1djBwaExaWk5xVDB6V3RqdVY3RnludmJobWJpZlMvYVZaMFZocGYvd0hnakhmTGEKdmR1eDBBdUg5YUV6bmtWSmN0TzhCcjdjZHp5N0tITkwyR0h1NVlFaVJKT0JmOTRYUnk1Tk9Ja0NRQVFUQVFJQQpLZ0liQXdJZUFRSVhnQVVMQ1FnSEF3VVZDZ2tJQ3dVV0FnTUJBQUlaQVFVQ1ZRblRQd1VKQ1oxcVZnQUtDUkFnCnR5ZzYvaVZNYWVyYkVBREdySDcrTFVpeHBxVlF5TGdxRWZwT3R6NFZWcHRHZlZDdUFRb2gwU2RkemJwNEZmRVkKMkk0cm1qRnJlVnRxc0NJZndkWlhEOWtGYUF0NzlzUmxXcTRMMlN4eks4Wmo5aVNPUjBNYjJxQzRjYnNkbmJZaQo1OEFHTVhFK2xjQVpiV01wVExMbmNIWHU2TldqMXI2YWlMVGNad3hZV0FwTGMwWlhOZDNQdHlIeklGdnR4RGZ2CmRqa1ZBYkJRc245U0NaUng4UWVjaVkzL01MNERvbCtSbkcvOEVLTUQvU3BPTFZPSmk2ckhuc3AvR2piR3V1dWkKMUY5Um5lZXd3a1VkM3MyaXg5YWkxdXZLK1VQc0dVanFLQ3NRVVJnN1VjM0N2ZFlWMDRVODRMTWxoUU1FUDBqLwpMUnNPb1BXR3pxQ1JxV2NKbDNhOTdvTkh0N2QvM05OTmNPUVg5OXQwZEF6dXFPSGpIR3VSSEdLYmZBTU0zUjErCjlCMm1kQ1E3ZnRmNWVWNktlSHJ6bytNbTNhSDh2RXo1NSs2Q3F1c1AyN2hpN21WaXdrWG8vcGZndXVRSlNENXcKeG05d0x3OHViRjZDdmFPQlp6UUNJc2EwczFCTmpjbjNqaGg3UmNRK0NNSFFiTkFHeUtkU2dyeWhOejl6enY1VgoxaFdGWGhiY1BxSEpwYlVMWmtLVFVJTUpvamc5dUFrTUhzMDNDM2pOUWdnTFliZWdBN2xkN3ExMlBubDAyMDdNCi9ZbDIrdzNiM0hIb3E3TmVSVHUyTVFtQ20vcWt1Vk9ZNlJ4OVFITGMxU2ZHQ1RaUGpLUmtxSXJnTVZOUjViNXoKNlVsV2M3MzVkMThEYmV5QThka1BwSUd5a1pnM21LR0VZTDUvT3hHU2ZJNldEK21rckF3UE1SVHZ0SWtDSEFRUQpBUW9BQmdVQ1VrYzJGd0FLQ1JDWE1sM1krZjNWQmtlZEVBQ2wra1A3SGJTWThRdmkxODRDbjQyS1htWTlrYkIzClkzSUxwcjE3VU9zT3lwZWY3Tk5qUWFudGR5c1dSZGNQc3JLRVFZQUtGcGRBd1B5MEVFVDZodXVFS25CYjRUMC8KL0tqYnBFbE9HaFBCbThYOStabGJwWkZ0N3pVS1IxWi9BSVYybU9ZamRpamQ5bEhqalF0MEwycWFxd1g1MEluQQpFVWlNTUlVUGJheXhDRFJwVkpuNFZaWUVnbGJQUWZJazBLVVBKMHlGRzdESTV5THgyQjMzZUY2TVNzNWR4QnZmCmdTay9Xck5ISmxvcHVmSURyTnM3eGtNWXdoK042Sm4wdkV5U3dKVjR4RWlsTzVTZzRtdGxwREhWa1UzTVAxVVEKR0JFcUhNb2NLSXRYelFsUEZpQnNHU0dEWG9UQWFTRUR6MEcwbHlnU3IvYTJ5UW1tZ2JTeEs3cnJGU0tzV09BQgoycWZqVUN3a1NlU25kcWVWcEZpT0M3ME5icmtINjNyeVhFTHRYSXk1RGluOXp0OFVwd3Zndzh4M0tNTzJPSkdVCnU2OGg1ZkQ2ZWlEZmgraHZtNTIvOC83ckk3SjJFemxaSUVnZENQbVJYYkV3ZmlqS1FNUU0zSXRRSDh3Ymc4bkYKQjBpRjVNZkd6d1NpTkRtZFRjTzUrSXg1QktzeGdyb01HSmNmQjdWZUlEQmxFLzd4aFVqNEovMHJ0L1d4d21YZApjVU9DV1dPRmFYcTlLV3NSTEUvVjM3NTVWRktlcmZZLzZOYW43RDYvM0tqc0hhdDVmSEpYdm1zTWRBSWhEREI5CnNoL0x4MFIycHFSV0lZNDYvQ1h0SForQ3lYNXZOS1kzZnRsdDVlNTArVmZvUDByV3NaZldWR0F2WVlNN0w2dmIKMDZ3anY3YVRZcHpCeW9rQ0hBUVFBUUlBQmdVQ1VrYzlGUUFLQ1JCN1dGc3dnSHdxaDZXTkQvOWp5cmthbUV6SwpTTE1JVmc5MFZHdTFwc1JXK1JlZ0I2Ung2ZUhoZUZtSG9TTWFzRE5IK0YwNjNvdW80dWhyaHVGcEdWQXNJN2o0ClQ0MUVKazZJaVB4ZHdiQWEzemxCQUR2b3o5cmZnTm9lb0F3c0FEUkpuN09PY0h5SFp4MGJLSVNkWld3TytWQnAKM3VuUTJGVFZINnhzdG12R1o4cTRvbnQ2UlRLRlpKMHhTd1kwQ1owUUdUN2czT3lrY2tLakRSNmhkOUxZS1A0UgpVc3Z0Z2Q1cmUvTWV3KzJlOFJ1YlFzWTFWM2FvalRPcm8xQXFXNkhoaDlQKzNmRGlMSUU3aUlONWZWMEMyeTFoCjdpZmI0aEh2aWl3MiswR3M1RzBQWGhPNGtpa3pHTnk5WSt1Ylc3TjlJZmU2QWh4WTQ1U1RxS3o2aGhSVU4rZ24KbFhEQ3V6TzVnOXlCb1FQZ0xqQXRYYU5tUkJKd0NiaXRIbGxyUHpBbDhUQXVtU0ZMem5oU3U1M3RhUTVTS0M1NgpUTmtoR2Vydlh0NklNUmtTbEV5bzM4endhZUN0dEVYNjNJOEJHTTFxaEgxeGNldDJTUlFJZzUyaW1sVEdBQjhECnQ0S1VXNzhjT1V3NlM5SXJCY3F0aFRueVZxQUVITXpkcGp3bG55aGU5SkgxZEs3VGYxTms2d0tlT1lob2pqbW8KS29KdVpLMHJHbEVZRkhCemYvayt3My84Y1d5L0ZsN3duZkZrQWZDZXZ3KzdKRUthb21uSFpLUmk1NFh6OFBPTgpiRzhxTWxqNTUxYWlqRW16T2tBandKUzVqbGlxeHZaL050SFc3emxwc3NPWnBQdG55d2tCbjVwaXBwTU1jSHNwCklnT3lLNFRBbG42Q2t5WXBxNDlhUm8zVXAxaTFFUmlCejRrQ0hBUVFBUWdBQmdVQ1VrYzZOQUFLQ1JDN3gxYmQKdmxsZmEra0ZELzk3TngyWjZlV1pFUXFBeHpKcHVyYmRNWmhOeEVQbTJhMnNhc1JvRkNTSXkxbjE2MWhoKzczRAp5Zy9KdU1uQzAwdS9VV3Y4WnZpMW9yUFV3bGRwUXRoeWF1SHhPdzg2VnZlTTdObTZleld5SkZIMDhKVWM2QmE2CmhlL2dERlV3MWYzMURXTkFPYkw2N2ppWFJhRFJhdFAyZTVFRGZLY2NpYnJSN0NXbTR5NjNXcEw1MXdyVFJCRjAKbU1vSjVYV2dFM2JES1N1S0V0SlF1eXpnTUNBejdVWUxwaUJHV2N0QTB5R1JqREN4OHJUUnI3SndJYjExb2lWSwo3Kzl2UkNLSWxzSHp3anNveDRtWE9ESFcvM3R1SHRRbmZvRlhjR1BPQlpkcnlGVmE4YzZaTVdoWFZ3Q1h3VWtOCjcreXhoNFBIY09lQXlYQ2JtOWNOd2hQMFQ4YW5BZmRFYVd0NFYvWXp6TE1UL0QvTGVyRHdBM3FwZVp0TjE3UTgKNkFNRzkxbzd2dUNBSnZOdTFNL0REMHJOOVpoSFh6QlkvRzhiMjFDZm1qVEpRV043QkNVTFBWWDFWdGFUUU9JaApnYTdqZEFPcDBVUlZKbHNsSFFHRW03SFkvTXFWNTIyZGU0S2VEMWNlU2ZuODVYNXVSd1NFdE9HeGwzVjVsSWM2CnA5a0J6bVpYekpsS0c0dDJldFBhRkkvSkxxQkxoUHpqbFk4RDRZSS9uUWFkNHNoWXYxMlBlUGNNTXpsOHdueGIKcUNtOEdsczdPYVAvYkVhR0JYMjhhREgvU2lRQTB6dUs0WXpJeVFneDluUjNRL3RDOGNiWmJEK1QrN3AwUjNyQgpHWnRDcXhUOWtBR3VySUlQaExvYTJuVGlWWDd0L3dSZjNYL2U0MmJoMGllTlNTZ2ZHMU9VejRrQ0hBUVFBUUlBCkJnVUNVa29ZMVFBS0NSQXNaR1N2S281TUF2aDhFQUNoODdGakN3RGR0TDVVc2dnTEoxalZlS25Ud0VINFlxK3UKZVRpUEZGdU5Oem9hOWswRXBtcURxVG5samtLTmJsT0F2blZ4cGF3VjljUkxIdFdTNGNYWGdRUWJWUUU2Q3hyeQpoQ0NzZmlTVjRtOFlLVCtSSzNLeFVCOHhKRW1oWVVYZjFBZGQybWM0MkVGVTkvSWF0TVdlUkgvNzdyWk5PZUNDCjVHQ0dWM2lVcDN1OGNqQi9xeVQ4V3ZBeUtTeTVabnIzbVBMZkhMeUZKdkZUbHA4OWhoc2xYTHhOL2hIUTlwdVcKZDhGSEFLYW5CWmtKSHgzWmJGbUs5UE8xOFExWGtycXBUUWNGSmIwT0kzYmNBTVZIN0pjZVNlSERrRjB3SFF3UQp6RlM4c2tubUYrVEdIR0l2Y1pQZWxDS2tDT0dXUXNoR3RRODV3dXY5ZTg1SlRjYlNqUlFoTkh2ZzBYNkJ5bG5mCjVKWFpRaDFyWkRaQjdvK0ZNR1ZVcExzamNuWmRFTGdWZEgwbUVLWEhLZkl4dGN2ODhXZktzWlFrVXVTZVNaK3cKTmZvd3pCTVZFNCsrazZHV21EVVEweEJJL2M0aDN3NjJ3dll0cFBoRzdOc1ZXdjlhL1o1cURuM2lmNzQvLzJveQpWQ2pDZUt1QVBpOUR0VWVaMDJvTjkxcTE1SjNIS2dOeFYyR0VLdC9wVVRwNEZ1SWVLUFpFUDlRMTg3S3Z0NHJBCmZ4NUxhMW5mZG9QUkQxaUV5SGFBbHNscWpXQk0wY1JDcWh6WmxmQUxYbm84UWRjZEM4dWxsMTJQOGZzcFUweUgKOHdhdzVvL2N3S0NuVjJ4bVV0dElMRk9iMVFGdzZ6V3lCWTBlN29GdWRodkcxaDRCdDBnMC81cVAyVFJDdnBoNAp5U3pibUVINVA0a0JJZ1FRQVFJQURBVUNWSjNCNVFVREFCSjFBQUFLQ1JDWEVMaWJ5bGV0Zk12WkIvOUlxN2lEClRYMXlES29GZGU2UkswaENlTkVwV09WOVAwR3VXY0dmVWQ5cGFJNktIbkx3dTQ5ejlLZ1dXMEhOMkpXK0h6TkoKNDB1NWsrQndvOENjL0RqZUJvRU9WOFBCR2tKckVqSW9mcng1NlFaam5uR2dCclRvdUcyVGFYL3pGamR6aXVmYwppQkpIRHhmaFJJTTFXNDdkSVBQTTFtMGY5VWlYekt4eVl2R0NEbWtKdGp6Z0FubDRGdms0TW05S1o0cVVBVHE4CnBiMGxiZEkzZjAycm5kZ0x2L2VWNER1OHRaMDc1TnZvZEE4UWRkd2tqRXpMNHFtSFdoNHJqMitvdk1PeThYcTgKdnVJN1BVR1VHVG8yN3hYL2tyTFVmeVNiR1NqR0RHc0JYRnlmWjZQUTB6UjF2eXFoZ0t3Z1B3M2VaenpsZ3k1YwpTL0Z6MmhmNjlPVHRoR1djaVFFaUJCQUJBZ0FNQlFKVXI0L0pCUU1BRW5VQUFBb0pFSmNRdUp2S1Y2MThFNWNICi9BODBmWEVvWWVwd3YrRGtrdW94WmdwR3ptQ29QZEROcmFuUHpuREs5ZlBLSlBCTXYyMFJDTTh4emtuU1FiR1oKcmY4ZVoyTUN5aHlRV0JKREFnU1B5dDJMZ3NtTjRJTlRCQ0xaZ1dqZnorTXFoQk9zTm5IUHVUbW16TEFTK3pCUgp1cWtRQkJEUkZKMTBxUkNwZ3RmZTlRajh3YS9YUlE5REt6MjRsYjhqL3pnUXp2Um5MUjI3K1RrVUc1aTM3VityClBjR2daeTlBcDV6V1NjQzUvTVBZZ3lBM09DaG81cDMvV29KMHB0cXZJaWMwbUJWUjBFV2ZTbkptbFAyaDF6OGcKTGMyVEFPZGZkWGY2eERQTVpLQmw4dEh2SnR2WFM0K2h2QVcybzlDeWhtcENKVkZsUDU0cDJzdktMQkVxZ1Y4agpqcVgxbDhObjRadGZVb1orR3dKYmFlcUpBajBFRXdFQ0FDY0NHd01GQ1FXam1vQUNIZ0VDRjRBRkFrL0dFWTRGCkN3a0lCd01GRlFvSkNBc0ZGZ0lEQVFBQUNna1FJTGNvT3Y0bFRHa2paQS84RDRZSzFQZjV2RVlGSmF4OEYzYUoKV3B2Z3Y1QVpXbExRa0lWZ1lLQUMrU3J3QzdvejRpOVZGRll1Q0FjT0NsUkFPMll3RFpITVVkd3Q1MXR0SFRoegoxUXUzUkVodk8wTFg3dDN4a0YzWkV4NkdVdEhRWkkrTzFJYk5JbXh6WWpvTkVMb1o3aEhraHh6dWlzbzUwOWRwCkVLaFEwSm5Kb0FIOTlrU1JYcWV3Y0xrRjNqTXR1eWczTDROeGQ0NVpBcExuMW5NRmtyWE8ySmlyU09GMUMzcVMKbEc3d0tobHEwTFFSYVE2ZzZMbU1YUm50ZlJpSnBwOVg4YVByN0gwWEh6SWNMM21TS2g2NWE3VklOV01yR1JDVAprbzQ5cy9JTnR3d2xQeU9ZL0UrTEI1cnJEWm9JdTdua3pibmJxSnZkWjN1ai9jNkFLNDVEdjFUeWRGTXJIVC8vCmJXWDI2MzhjSzR4RkVkM0cxU0dHZy9zWU1NVzJTYWtqUk9NbTFaNnJyd3U4YkRPRlFsLzl6YnU1OHRraG5tbFQKQkpDQjhPNGMrQXhDM0pqaDlkalk1YnU5NkxOQlVRY3k5d1ZYb0dyR3FkWDFQdWVTUnM1MFRqNDFDQktaNjA1UQpiaDhtSWVjVEpRNjJ4L2ljUldpNHl6eWRGVHZkbCtkakg5eUhEUXUxbEptNE95MHB1b3BUMzhwYnpBUUhmZzdpCjAxQXB3dThBekc0dVFFcHRTR2pyT2R1MzAzNW9DbldpTXF0dGlWOFdDalFlTENCL1llNi9WVTY1YlZ1TG1kOHQKbnlRRmFEYlVCMDN2K29UYXVOTUhyQXJuOHk5MHBXM0JGbVdQMHpkeWRIYlVZVm5WV0d6NkdjdGxtc2RrRXF6Uwpzd2d4bGY0enBhM2p0dDU3Y2c1OFdvKzBKMHRsYVhSb0lGZHBibk4wWldsdUlEeHJaV2wwYUhkQVkzTXVjM1JoCmJtWnZjbVF1WldSMVBva0NQZ1FUQVFJQUtBSWJBd1lMQ1FnSEF3SUdGUWdDQ1FvTEJCWUNBd0VDSGdFQ0Y0QUYKQWxVSjB6OEZDUW1kYWxZQUNna1FJTGNvT3Y0bFRHazRyZy8vV0RkNnpuZmV6TVdjOURDVU50NGxEdWRvWmQ1cApQUE1FU3lFYlB2b0pvYmZlSHVXTFNDNGpFWUYybk1lYWIrc1RhUEQwV3RSd2J5cEw1NzJzUkFPMFlWUHBSU2VJClNzZVF0NCttcFMrN0FyTFhsMHdiN296MXJYTFcyTnNnVGxDY240VytVcE51Y3FsUjgyRHJNMEdaa2ZVUjVlTDUKb3d6UnpKVDVKcndtU1MycUpoT2NqTnkyVjl2ZEJJVU9mNW5sYmJFREUxMmJsNlpmMzhubG9Yd29FaThXTGJrSApYc3BORkVpSzZqS1lJTkJZdlJLYm56Um5LbHdRQWRDTWVESzFFYXlvbkQvMDltN3I3Vm0wdXVYT09ZMDNIMENYClU0VzBtY2h4S1ZKbGR5TDhiQ29sRmZxQ1dmdHFraHBTVnJsMEhaa2h1ZjhiY0RWcjhaaE9VSmxGeVVTQnNNbEEKRnVXZUlXdHdLQ1ZNaXpvakRrLzEwQzR5YndRRUZ6RDAvMHI2L2pQNjJqajIyZm9DVC9BV3QvNk9TczJjTk1jegp1Uy9SdnpsWmhJZE1rUkNTQ3dZdlZwVW5mamNkcG5zU0hWczl6ekxXUXBaUEY5djYrMm1pckFsUS9hZ3JTanVVCndINHNScXlBNkJtQU94dnNXMEIzek85eFF5Rk9QVGY1dXNYclk4UUFkTDQ3cGtBZkdZK0M2ZStRYk9BRjluWnEKUHBwaHprSGZEWm9yZWxCSTZoK0d5T0wzNHVNTkRodVdHSDRxNXFGOW4vN25Rb3dLcCtjVXJ5YlN5KzFSendTRApRdFF0L1RYSStyRnQxcUUzNkM2ZGJNek1VUGoyL3pPMG1GQlV1Qk9YOXJmZE9RYlN4SjJRWVM1ei9adnAvM1N4Cm5zVDFWZW54YkVBS0d6K0pBU0lFRUFFQ0FBd0ZBbFNkd2VVRkF3QVNkUUFBQ2drUWx4QzRtOHBYclh4Qmt3Zi8KVjdKSTRoUkxmNTF6Znl4Q1BqQWcrOGxCcFhPZU9sS05XWFZYdTNVdW44cmtOTlpsaVFzRU4ybEdqdDBSc0d3Qgp3eHAzLzFIWmVNLzlqczdUUWtBVUNLZmMyM3dncDlyOHNHREVZTm9NUTRlNktCN2VkMGhRWmJadFFPRzZNS1ZBCnFxNHN0dlFwajh0QUk0Q29JNTc3VXNkcS9ONlh4bit6OUlkVnFNTGRQL0REV1dDL21Felpnb0FKb1IreGp1ZGkKUmozcVFXellVbVdVTU1BWitOeW9QK2UvV2I4S2Y1d2lkaDRDcUJodko3ajJaSWw0NFRtUUdUN3Y3UjBrVWkwVQp2WEJkQmRoQnIzeVpLZHdSdXpQdE1CN1RXbkZGN2h6TzhlRk5yVG5Nc2VsUzFnd2lQODRIWFh4ck1YeXpLZ3VhCi9ldjY5cURlWENUSmkxbWk3Y3Bmb0lrQklnUVFBUUlBREFVQ1ZLK1B5UVVEQUJKMUFBQUtDUkNYRUxpYnlsZXQKZkVwd0Ivc0VKbGlvczNNVndRV3lCNkc3ems3RThZR0hNSHpETnZ1dzBibldRdkc4Z0Q0RVUwbFA2dC90RndRaAp3RHBzVXRFczBHVzVieVNIM0VaYXdCODZ0VjM5Mmg5WHk2VE1xU0w4OEZRL0FtS2kvK0VBVWs1c001Yk8ySUJOCjVnSjJjY1RMelhGL1NOamVJV1ZpeUR3dXJZNlRPVUliSzRyekltVFBPTnc1SER2ejdmRzdSODFFMnR4ZFNVU20Ka2pOVWVYSk5FNVhzaU5TRUp3cmpjMHk4bTdWU0lRTGwycEJTcW1hVTl4VW5WSndxd0tSVXF1ek1CSUljQk1oaApSMFNRYndZMERKUXhENHpKRUQvdElWUTVpaFZuNVN0Mm1FbTIvZCtRTEozTmRkajJ1UFJEdG00cmRmQkIzanhQCk1sZy93dVdiZ010a3dzYUN5UUFpR0FUYk9tN2R1UUlOQkU4dXorWUJFQUN1WFhnYVlHanQycDI5a1NYMEhtenEKY2RJb2JLNlcxUE5Ob21YQ2FJRG45MkxKR0N0K3pzeUs5RzBaQmt0cEEzNWF0bjEvVkxCUUZ6cmJQd00zR29VQgpuZGprWHFHQ0JzWExiTzRYYTY2U1hGTHJKTlhEd29yNmtKTWZsNVovaXM4MHNybE43dEZ0NUZGUERxTEplUzllCitCVy92T1dIN2dqZXR4akRjc2M5QlVZNGtqOVl2bU10ZTdOT3hLUGhjTytsQ2R4Zkg0QVZ4ZmNzT21HQ0JPS1IKdko4eXA5OXlLb3Y0WDI1RWRhNjNydnVMV2tXNzhWVWlZWG5HcnczNDRNTkI4UWVoWVhVMlJscjJEUnFUU2JyQQovSmQvZzJZSmRFMThSQ1RJSTlNbm5reVQzbk81SDE2eDVDeE1tWXE4L0V5QkdaZWRQYVhqKzdGZDY2WUl4QUplCjh1Q1ppbHFOU0hpNXgwOGFSdCtEQ255UDJqN3F0Y3ZUUkVyMGRvT29VNDBOajdqRFphT2ovOEVxbktNUTlDMXIKZ1JMWU9xV2dTY3ZNN1hPYVIrMHBIbHN0Zk94RXdNcVdRbG1Vd3NnZ21XVFduNmhueWY5c3NWRFl6WkRJa05KMApkaEVkbk1OQVl6bDVQVFBldEtOaFhiMmhZVS8zYXA0emJkeVVwS2VGWWNTRG5qSDc4QXg4VGFFdXJneE4rd3ppCjJESFNRTVRUU0RIVDFqM3YzL1V5TklaUU04SnVhZml2aitXUXlNR05nUlQrYWJpL3gyRm9jeDFEL050eWZOM2cKMTFzdW9iWkxnYkplWVpmZGZ3TzcyMTRqRGlTcXJGRFVzUHpRL280djdwSEY1NVpYSUVuYnJXeWdaSk9qelRwMwpIWHRNekF1bG5qYmJzWTE0amlJbWd3QVJBUUFCaVFJbEJCZ0JBZ0FQQWhzTUJRSlZDZE11QlFrSm5XbzhBQW9KCkVDQzNLRHIrSlV4cFYyb1ArUUU0WU5QeThRZlJUbUV3c0VZMGZUaW4vTjR4Rm0rSG92U3ZKbUp0RW50bDZzb0UKWXFLRDltOTl6ZEpCeDB4cjdTNEV1QkNkRjJFM281bjlYbjFQNXA3OXN2bGttamVkYjBJaTJudGFRS1hnSmVFRgp6amJnZHdzY0dKZy9jRVdwaGs0V2xzR2FqVFRrdld0MnlvanZMRUdOaU1RZmd2V1gzbytXU0VRaTJ3YklnS01hCkI3clJiTGdMQ0xraXBVZVB1dDZkcklaeVRYNGdRazZBWnZGMGloUjJBNWpzYzY3NnQ2OXpiU0hPdmxGOVY3NEUKMFYyaXNsZzJPWTl6TDkzVzlEKzh4bmdkdmE5RVlWdzljR3VuazNITGUvdHRVSUMzNTFTNGRhU1V2MlRrTGdmegpzWXFLL0syUkFuYmxqMStncWM3WHFFYXh2b0d1V3ljMmp2SyszWFBOVTQwdEtaV3pGT0hwMTRpNWFlL1JmQzAvCnppQ1pKK3JEdjlaNFpSMERvTXNiZmtqcGs1b254MHAveU5WVk1NYmdkanNubm1UMnZOWDgzNmpOUEg0SmZidjkKbSt5Skx5V0FGY2FyRU9STHRTN0FWeG1BVzRPRm9uZnRjQ3NLaGtIRWpSb2VINVJmbGs0VVJLNzkveG42b3hLaQpJYzh4emt0SWtqWEpVbFVENjJYZ1BwcW5XTXZlT1NrVGlGa21oUG5kZmNOVURwOCtleVc0dEY2TTNJcDBML3A3ClU3SDNJK3NQNnhQVlV4NXB2VUZRdFFoNWJmdWt4TStQc3RvQ09oNSt4TW5mc2dLNHg2RFAzSWE4VmtxRm8yckUKKzJFWDROZ3RyK2lvbzlRbS9qSDZGenlOZXJ6RjlwU2RsZzVmK25xUXNrSHljVlJxdituMU11RS82czA4Cj0zY1VSCi0tLS0tRU5EIFBHUCBQVUJMSUMgS0VZIEJMT0NLLS0tLS0K"
	if key, err := base64.StdEncoding.DecodeString(key64); err != nil {
		t.Fatal(err)
	} else if _, err := ReadOneKeyFromString(string(key)); err != nil {
		t.Fatal(err)
	}
}

