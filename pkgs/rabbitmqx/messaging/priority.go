package messaging

import (
	"strconv"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Priority uint8

// MessagePriorityKey 消息优先级元数据键
// RabbitMQ 使用此键从消息元数据中读取优先级
const MessagePriorityKey = "x-priority"

// 消息优先级常量（0-255）
// 建议使用常用的优先级常量，如 Priority0, Priority10, Priority50 等
const (
	Priority0   Priority = 0
	Priority1   Priority = 1
	Priority2   Priority = 2
	Priority3   Priority = 3
	Priority4   Priority = 4
	Priority5   Priority = 5
	Priority6   Priority = 6
	Priority7   Priority = 7
	Priority8   Priority = 8
	Priority9   Priority = 9
	Priority10  Priority = 10
	Priority11  Priority = 11
	Priority12  Priority = 12
	Priority13  Priority = 13
	Priority14  Priority = 14
	Priority15  Priority = 15
	Priority16  Priority = 16
	Priority17  Priority = 17
	Priority18  Priority = 18
	Priority19  Priority = 19
	Priority20  Priority = 20
	Priority21  Priority = 21
	Priority22  Priority = 22
	Priority23  Priority = 23
	Priority24  Priority = 24
	Priority25  Priority = 25
	Priority26  Priority = 26
	Priority27  Priority = 27
	Priority28  Priority = 28
	Priority29  Priority = 29
	Priority30  Priority = 30
	Priority31  Priority = 31
	Priority32  Priority = 32
	Priority33  Priority = 33
	Priority34  Priority = 34
	Priority35  Priority = 35
	Priority36  Priority = 36
	Priority37  Priority = 37
	Priority38  Priority = 38
	Priority39  Priority = 39
	Priority40  Priority = 40
	Priority41  Priority = 41
	Priority42  Priority = 42
	Priority43  Priority = 43
	Priority44  Priority = 44
	Priority45  Priority = 45
	Priority46  Priority = 46
	Priority47  Priority = 47
	Priority48  Priority = 48
	Priority49  Priority = 49
	Priority50  Priority = 50
	Priority51  Priority = 51
	Priority52  Priority = 52
	Priority53  Priority = 53
	Priority54  Priority = 54
	Priority55  Priority = 55
	Priority56  Priority = 56
	Priority57  Priority = 57
	Priority58  Priority = 58
	Priority59  Priority = 59
	Priority60  Priority = 60
	Priority61  Priority = 61
	Priority62  Priority = 62
	Priority63  Priority = 63
	Priority64  Priority = 64
	Priority65  Priority = 65
	Priority66  Priority = 66
	Priority67  Priority = 67
	Priority68  Priority = 68
	Priority69  Priority = 69
	Priority70  Priority = 70
	Priority71  Priority = 71
	Priority72  Priority = 72
	Priority73  Priority = 73
	Priority74  Priority = 74
	Priority75  Priority = 75
	Priority76  Priority = 76
	Priority77  Priority = 77
	Priority78  Priority = 78
	Priority79  Priority = 79
	Priority80  Priority = 80
	Priority81  Priority = 81
	Priority82  Priority = 82
	Priority83  Priority = 83
	Priority84  Priority = 84
	Priority85  Priority = 85
	Priority86  Priority = 86
	Priority87  Priority = 87
	Priority88  Priority = 88
	Priority89  Priority = 89
	Priority90  Priority = 90
	Priority91  Priority = 91
	Priority92  Priority = 92
	Priority93  Priority = 93
	Priority94  Priority = 94
	Priority95  Priority = 95
	Priority96  Priority = 96
	Priority97  Priority = 97
	Priority98  Priority = 98
	Priority99  Priority = 99
	Priority100 Priority = 100
	Priority101 Priority = 101
	Priority102 Priority = 102
	Priority103 Priority = 103
	Priority104 Priority = 104
	Priority105 Priority = 105
	Priority106 Priority = 106
	Priority107 Priority = 107
	Priority108 Priority = 108
	Priority109 Priority = 109
	Priority110 Priority = 110
	Priority111 Priority = 111
	Priority112 Priority = 112
	Priority113 Priority = 113
	Priority114 Priority = 114
	Priority115 Priority = 115
	Priority116 Priority = 116
	Priority117 Priority = 117
	Priority118 Priority = 118
	Priority119 Priority = 119
	Priority120 Priority = 120
	Priority121 Priority = 121
	Priority122 Priority = 122
	Priority123 Priority = 123
	Priority124 Priority = 124
	Priority125 Priority = 125
	Priority126 Priority = 126
	Priority127 Priority = 127
	Priority128 Priority = 128
	Priority129 Priority = 129
	Priority130 Priority = 130
	Priority131 Priority = 131
	Priority132 Priority = 132
	Priority133 Priority = 133
	Priority134 Priority = 134
	Priority135 Priority = 135
	Priority136 Priority = 136
	Priority137 Priority = 137
	Priority138 Priority = 138
	Priority139 Priority = 139
	Priority140 Priority = 140
	Priority141 Priority = 141
	Priority142 Priority = 142
	Priority143 Priority = 143
	Priority144 Priority = 144
	Priority145 Priority = 145
	Priority146 Priority = 146
	Priority147 Priority = 147
	Priority148 Priority = 148
	Priority149 Priority = 149
	Priority150 Priority = 150
	Priority151 Priority = 151
	Priority152 Priority = 152
	Priority153 Priority = 153
	Priority154 Priority = 154
	Priority155 Priority = 155
	Priority156 Priority = 156
	Priority157 Priority = 157
	Priority158 Priority = 158
	Priority159 Priority = 159
	Priority160 Priority = 160
	Priority161 Priority = 161
	Priority162 Priority = 162
	Priority163 Priority = 163
	Priority164 Priority = 164
	Priority165 Priority = 165
	Priority166 Priority = 166
	Priority167 Priority = 167
	Priority168 Priority = 168
	Priority169 Priority = 169
	Priority170 Priority = 170
	Priority171 Priority = 171
	Priority172 Priority = 172
	Priority173 Priority = 173
	Priority174 Priority = 174
	Priority175 Priority = 175
	Priority176 Priority = 176
	Priority177 Priority = 177
	Priority178 Priority = 178
	Priority179 Priority = 179
	Priority180 Priority = 180
	Priority181 Priority = 181
	Priority182 Priority = 182
	Priority183 Priority = 183
	Priority184 Priority = 184
	Priority185 Priority = 185
	Priority186 Priority = 186
	Priority187 Priority = 187
	Priority188 Priority = 188
	Priority189 Priority = 189
	Priority190 Priority = 190
	Priority191 Priority = 191
	Priority192 Priority = 192
	Priority193 Priority = 193
	Priority194 Priority = 194
	Priority195 Priority = 195
	Priority196 Priority = 196
	Priority197 Priority = 197
	Priority198 Priority = 198
	Priority199 Priority = 199
	Priority200 Priority = 200
	Priority201 Priority = 201
	Priority202 Priority = 202
	Priority203 Priority = 203
	Priority204 Priority = 204
	Priority205 Priority = 205
	Priority206 Priority = 206
	Priority207 Priority = 207
	Priority208 Priority = 208
	Priority209 Priority = 209
	Priority210 Priority = 210
	Priority211 Priority = 211
	Priority212 Priority = 212
	Priority213 Priority = 213
	Priority214 Priority = 214
	Priority215 Priority = 215
	Priority216 Priority = 216
	Priority217 Priority = 217
	Priority218 Priority = 218
	Priority219 Priority = 219
	Priority220 Priority = 220
	Priority221 Priority = 221
	Priority222 Priority = 222
	Priority223 Priority = 223
	Priority224 Priority = 224
	Priority225 Priority = 225
	Priority226 Priority = 226
	Priority227 Priority = 227
	Priority228 Priority = 228
	Priority229 Priority = 229
	Priority230 Priority = 230
	Priority231 Priority = 231
	Priority232 Priority = 232
	Priority233 Priority = 233
	Priority234 Priority = 234
	Priority235 Priority = 235
	Priority236 Priority = 236
	Priority237 Priority = 237
	Priority238 Priority = 238
	Priority239 Priority = 239
	Priority240 Priority = 240
	Priority241 Priority = 241
	Priority242 Priority = 242
	Priority243 Priority = 243
	Priority244 Priority = 244
	Priority245 Priority = 245
	Priority246 Priority = 246
	Priority247 Priority = 247
	Priority248 Priority = 248
	Priority249 Priority = 249
	Priority250 Priority = 250
	Priority251 Priority = 251
	Priority252 Priority = 252
	Priority253 Priority = 253
	Priority254 Priority = 254
	Priority255 Priority = 255
)

// 常用的优先级常量别名（便于使用）
const (
	PriorityUrgent   = Priority255 // 最高优先级（紧急）
	PriorityVeryHigh = Priority220 // 非常高优先级
	PriorityHigh     = Priority200 // 高优先级
	PriorityMedium   = Priority150 // 中等优先级
	PriorityNormal   = Priority100 // 普通优先级
	PriorityLow      = Priority50  // 低优先级
	PriorityVeryLow  = Priority25  // 非常低优先级
	PriorityLowest   = Priority0   // 最低优先级
)

// WithPriority 设置消息优先级（RabbitMQ 特有功能）
// 优先级范围：0-255，数值越大优先级越高
//
// 参数:
//   - priority: 消息优先级（0-255）
//
// 返回:
//   - PublishOption: 发布选项函数
//
// 示例:
//
//	messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithPriority(10),  // 设置优先级为 10
//	)
//
// 注意:
//   - 队列必须配置 max_priority > 0 才能使用优先级
//   - 如果队列未配置优先级，此设置会被忽略
func WithPriority(priority Priority) PublishOption {
	return func(m *message.Message) {
		// 将优先级存储到消息元数据中
		// watermill-amqp 会从元数据中读取并设置到 amqp091.Publishing.Priority
		m.Metadata.Set(MessagePriorityKey, strconv.Itoa(int(priority)))
	}
}
