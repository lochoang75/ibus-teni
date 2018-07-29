/*
 * Teni-IME - A Vietnamese Input method editor
 * Copyright (C) 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com>
 * This file is part of Teni-IME.
 *
 *  Teni-IME is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Teni-IME is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Teni-IME.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package teni

type CR struct {
	C rune //change to C
	R bool //is revertMode
}

var (
	r_a = &CR{'a', true}
	r_ă = &CR{'ă', true}
	r_â = &CR{'â', true}
	r_e = &CR{'e', true}
	r_ê = &CR{'ê', true}
	r_i = &CR{'i', true}
	r_o = &CR{'o', true}
	r_ô = &CR{'ô', true}
	r_ơ = &CR{'ơ', true}
	r_u = &CR{'u', true}
	r_ư = &CR{'ư', true}
	r_y = &CR{'y', true}

	c_á = &CR{'á', false}
	c_ắ = &CR{'ắ', false}
	c_ấ = &CR{'ấ', false}
	c_é = &CR{'é', false}
	c_ế = &CR{'ế', false}
	c_í = &CR{'í', false}
	c_ó = &CR{'ó', false}
	c_ố = &CR{'ố', false}
	c_ớ = &CR{'ớ', false}
	c_ú = &CR{'ú', false}
	c_ứ = &CR{'ứ', false}
	c_ý = &CR{'ý', false}

	c_à = &CR{'à', false}
	c_ằ = &CR{'ằ', false}
	c_ầ = &CR{'ầ', false}
	c_è = &CR{'è', false}
	c_ề = &CR{'ề', false}
	c_ì = &CR{'ì', false}
	c_ò = &CR{'ò', false}
	c_ồ = &CR{'ồ', false}
	c_ờ = &CR{'ờ', false}
	c_ù = &CR{'ù', false}
	c_ừ = &CR{'ừ', false}
	c_ỳ = &CR{'ỳ', false}

	c_ả = &CR{'ả', false}
	c_ẳ = &CR{'ẳ', false}
	c_ẩ = &CR{'ẩ', false}
	c_ẻ = &CR{'ẻ', false}
	c_ể = &CR{'ể', false}
	c_ỉ = &CR{'ỉ', false}
	c_ỏ = &CR{'ỏ', false}
	c_ổ = &CR{'ổ', false}
	c_ở = &CR{'ở', false}
	c_ủ = &CR{'ủ', false}
	c_ử = &CR{'ử', false}
	c_ỷ = &CR{'ỷ', false}

	c_ã = &CR{'ã', false}
	c_ẵ = &CR{'ẵ', false}
	c_ẫ = &CR{'ẫ', false}
	c_ẽ = &CR{'ẽ', false}
	c_ễ = &CR{'ễ', false}
	c_ĩ = &CR{'ĩ', false}
	c_õ = &CR{'õ', false}
	c_ỗ = &CR{'ỗ', false}
	c_ỡ = &CR{'ỡ', false}
	c_ũ = &CR{'ũ', false}
	c_ữ = &CR{'ữ', false}
	c_ỹ = &CR{'ỹ', false}

	c_ạ = &CR{'ạ', false}
	c_ặ = &CR{'ặ', false}
	c_ậ = &CR{'ậ', false}
	c_ẹ = &CR{'ẹ', false}
	c_ệ = &CR{'ệ', false}
	c_ị = &CR{'ị', false}
	c_ọ = &CR{'ọ', false}
	c_ộ = &CR{'ộ', false}
	c_ợ = &CR{'ợ', false}
	c_ụ = &CR{'ụ', false}
	c_ự = &CR{'ự', false}
	c_ỵ = &CR{'ỵ', false}

	r_d = &CR{'d', true}
	c_đ = &CR{'đ', false}

	c_a = &CR{'a', false}
	c_ă = &CR{'ă', false}
	c_â = &CR{'â', false}
	c_e = &CR{'e', false}
	c_ê = &CR{'ê', false}
	c_i = &CR{'i', false}
	c_o = &CR{'o', false}
	c_ô = &CR{'ô', false}
	c_u = &CR{'u', false}
	c_ơ = &CR{'ơ', false}
	c_ư = &CR{'ư', false}
	c_y = &CR{'y', false}

	r_á = &CR{'á', true}
	r_à = &CR{'à', true}
	r_ả = &CR{'ả', true}
	r_ã = &CR{'ã', true}
	r_ạ = &CR{'ạ', true}
	r_é = &CR{'é', true}
	r_è = &CR{'è', true}
	r_ẻ = &CR{'ẻ', true}
	r_ẽ = &CR{'ẽ', true}
	r_ẹ = &CR{'ẹ', true}
	r_ó = &CR{'ó', true}
	r_ò = &CR{'ò', true}
	r_ỏ = &CR{'ỏ', true}
	r_õ = &CR{'õ', true}
	r_ọ = &CR{'ọ', true}
	r_ú = &CR{'ú', true}
	r_ù = &CR{'ù', true}
	r_ủ = &CR{'ủ', true}
	r_ũ = &CR{'ũ', true}
	r_ụ = &CR{'ụ', true}
)

var replaceCharMap = map[rune]map[rune]*CR{
	's': {
		'a': c_á,
		'ă': c_ắ,
		'â': c_ấ,
		'e': c_é,
		'ê': c_ế,
		'i': c_í,
		'o': c_ó,
		'ô': c_ố,
		'ơ': c_ớ,
		'u': c_ú,
		'ư': c_ứ,
		'y': c_ý,

		'á': r_a,
		'ắ': r_ă,
		'ấ': r_â,
		'é': r_e,
		'ế': r_ê,
		'í': r_i,
		'ó': r_o,
		'ố': r_ô,
		'ớ': r_ơ,
		'ú': r_u,
		'ứ': r_ư,
		'ý': r_y,

		'à': c_á,
		'ằ': c_ắ,
		'ầ': c_ấ,
		'è': c_é,
		'ề': c_ế,
		'ì': c_í,
		'ò': c_ó,
		'ồ': c_ố,
		'ờ': c_ớ,
		'ù': c_ú,
		'ừ': c_ứ,
		'ỳ': c_ý,

		'ả': c_á,
		'ẳ': c_ắ,
		'ẩ': c_ấ,
		'ẻ': c_é,
		'ể': c_ế,
		'ỉ': c_í,
		'ỏ': c_ó,
		'ổ': c_ố,
		'ở': c_ớ,
		'ủ': c_ú,
		'ử': c_ứ,
		'ỷ': c_ý,

		'ã': c_á,
		'ẵ': c_ắ,
		'ẫ': c_ấ,
		'ẽ': c_é,
		'ễ': c_ế,
		'ĩ': c_í,
		'õ': c_ó,
		'ỗ': c_ố,
		'ỡ': c_ớ,
		'ũ': c_ú,
		'ữ': c_ứ,
		'ỹ': c_ý,

		'ạ': c_á,
		'ặ': c_ắ,
		'ậ': c_ấ,
		'ẹ': c_é,
		'ệ': c_ế,
		'ị': c_í,
		'ọ': c_ó,
		'ộ': c_ố,
		'ợ': c_ớ,
		'ụ': c_ú,
		'ự': c_ứ,
		'ỵ': c_ý,
	},
	'f': {
		'a': c_à,
		'ă': c_ằ,
		'â': c_ầ,
		'e': c_è,
		'ê': c_ề,
		'i': c_ì,
		'o': c_ò,
		'ô': c_ồ,
		'ơ': c_ờ,
		'u': c_ù,
		'ư': c_ừ,
		'y': c_ỳ,

		'á': c_à,
		'ắ': c_ằ,
		'ấ': c_ầ,
		'é': c_è,
		'ế': c_ề,
		'í': c_ì,
		'ó': c_ò,
		'ố': c_ồ,
		'ớ': c_ờ,
		'ú': c_ù,
		'ứ': c_ừ,
		'ý': c_ỳ,

		'à': r_a,
		'ằ': r_ă,
		'ầ': r_â,
		'è': r_e,
		'ề': r_ê,
		'ì': r_i,
		'ò': r_o,
		'ồ': r_ô,
		'ờ': r_ơ,
		'ù': r_u,
		'ừ': r_ư,
		'ỳ': r_y,

		'ả': c_à,
		'ẳ': c_ằ,
		'ẩ': c_ầ,
		'ẻ': c_è,
		'ể': c_ề,
		'ỉ': c_ì,
		'ỏ': c_ò,
		'ổ': c_ồ,
		'ở': c_ờ,
		'ủ': c_ù,
		'ử': c_ừ,
		'ỷ': c_ỳ,

		'ã': c_à,
		'ẵ': c_ằ,
		'ẫ': c_ầ,
		'ẽ': c_è,
		'ễ': c_ề,
		'ĩ': c_ì,
		'õ': c_ò,
		'ỗ': c_ồ,
		'ỡ': c_ờ,
		'ũ': c_ù,
		'ữ': c_ừ,
		'ỹ': c_ỳ,

		'ạ': c_à,
		'ặ': c_ằ,
		'ậ': c_ầ,
		'ẹ': c_è,
		'ệ': c_ề,
		'ị': c_ì,
		'ọ': c_ò,
		'ộ': c_ồ,
		'ợ': c_ờ,
		'ụ': c_ù,
		'ự': c_ừ,
		'ỵ': c_ỳ,
	},
	'r': {
		'a': c_ả,
		'ă': c_ẳ,
		'â': c_ẩ,
		'e': c_ẻ,
		'ê': c_ể,
		'i': c_ỉ,
		'o': c_ỏ,
		'ô': c_ổ,
		'ơ': c_ở,
		'u': c_ủ,
		'ư': c_ử,
		'y': c_ỷ,

		'á': c_ả,
		'ắ': c_ẳ,
		'ấ': c_ẩ,
		'é': c_ẻ,
		'ế': c_ể,
		'í': c_ỉ,
		'ó': c_ỏ,
		'ố': c_ổ,
		'ớ': c_ở,
		'ú': c_ủ,
		'ứ': c_ử,
		'ý': c_ỷ,

		'à': c_ả,
		'ằ': c_ẳ,
		'ầ': c_ẩ,
		'è': c_ẻ,
		'ề': c_ể,
		'ì': c_ỉ,
		'ò': c_ỏ,
		'ồ': c_ổ,
		'ờ': c_ở,
		'ù': c_ủ,
		'ừ': c_ử,
		'ỳ': c_ỷ,

		'ả': r_a,
		'ẳ': r_ă,
		'ẩ': r_â,
		'ẻ': r_e,
		'ể': r_ê,
		'ỉ': r_i,
		'ỏ': r_o,
		'ổ': r_ô,
		'ở': r_ơ,
		'ủ': r_u,
		'ử': r_ư,
		'ỷ': r_y,

		'ã': c_ả,
		'ẵ': c_ẳ,
		'ẫ': c_ẩ,
		'ẽ': c_ẻ,
		'ễ': c_ể,
		'ĩ': c_ỉ,
		'õ': c_ỏ,
		'ỗ': c_ổ,
		'ỡ': c_ở,
		'ũ': c_ủ,
		'ữ': c_ử,
		'ỹ': c_ỷ,

		'ạ': c_ả,
		'ặ': c_ẳ,
		'ậ': c_ẩ,
		'ẹ': c_ẻ,
		'ệ': c_ể,
		'ị': c_ỉ,
		'ọ': c_ỏ,
		'ộ': c_ổ,
		'ợ': c_ở,
		'ụ': c_ủ,
		'ự': c_ử,
		'ỵ': c_ỷ,
	},
	'x': {
		'a': c_ã,
		'ă': c_ẵ,
		'â': c_ẫ,
		'e': c_ẽ,
		'ê': c_ễ,
		'i': c_ĩ,
		'o': c_õ,
		'ô': c_ỗ,
		'ơ': c_ỡ,
		'u': c_ũ,
		'ư': c_ữ,
		'y': c_ỹ,

		'á': c_ã,
		'ắ': c_ẵ,
		'ấ': c_ẫ,
		'é': c_ẽ,
		'ế': c_ễ,
		'í': c_ĩ,
		'ó': c_õ,
		'ố': c_ỗ,
		'ớ': c_ỡ,
		'ú': c_ũ,
		'ứ': c_ữ,
		'ý': c_ỹ,

		'à': c_ã,
		'ằ': c_ẵ,
		'ầ': c_ẫ,
		'è': c_ẽ,
		'ề': c_ễ,
		'ì': c_ĩ,
		'ò': c_õ,
		'ồ': c_ỗ,
		'ờ': c_ỡ,
		'ù': c_ũ,
		'ừ': c_ữ,
		'ỳ': c_ỹ,

		'ả': c_ã,
		'ẳ': c_ẵ,
		'ẩ': c_ẫ,
		'ẻ': c_ẽ,
		'ể': c_ễ,
		'ỉ': c_ĩ,
		'ỏ': c_õ,
		'ổ': c_ỗ,
		'ở': c_ỡ,
		'ủ': c_ũ,
		'ử': c_ữ,
		'ỷ': c_ỹ,

		'ã': r_a,
		'ẵ': r_ă,
		'ẫ': r_â,
		'ẽ': r_e,
		'ễ': r_ê,
		'ĩ': r_i,
		'õ': r_o,
		'ỗ': r_ô,
		'ỡ': r_ơ,
		'ũ': r_u,
		'ữ': r_ư,
		'ỹ': r_y,

		'ạ': c_ã,
		'ặ': c_ẵ,
		'ậ': c_ẫ,
		'ẹ': c_ẽ,
		'ệ': c_ễ,
		'ị': c_ĩ,
		'ọ': c_õ,
		'ộ': c_ỗ,
		'ợ': c_ỡ,
		'ụ': c_ũ,
		'ự': c_ữ,
		'ỵ': c_ỹ,
	},
	'j': {
		'a': c_ạ,
		'ă': c_ặ,
		'â': c_ậ,
		'e': c_ẹ,
		'ê': c_ệ,
		'i': c_ị,
		'o': c_ọ,
		'ô': c_ộ,
		'ơ': c_ợ,
		'u': c_ụ,
		'ư': c_ự,
		'y': c_ỵ,

		'á': c_ạ,
		'ắ': c_ặ,
		'ấ': c_ậ,
		'é': c_ẹ,
		'ế': c_ệ,
		'í': c_ị,
		'ó': c_ọ,
		'ố': c_ộ,
		'ớ': c_ợ,
		'ú': c_ụ,
		'ứ': c_ự,
		'ý': c_ỵ,

		'à': c_ạ,
		'ằ': c_ặ,
		'ầ': c_ậ,
		'è': c_ẹ,
		'ề': c_ệ,
		'ì': c_ị,
		'ò': c_ọ,
		'ồ': c_ộ,
		'ờ': c_ợ,
		'ù': c_ụ,
		'ừ': c_ự,
		'ỳ': c_ỵ,

		'ả': c_ạ,
		'ẳ': c_ặ,
		'ẩ': c_ậ,
		'ẻ': c_ẹ,
		'ể': c_ệ,
		'ỉ': c_ị,
		'ỏ': c_ọ,
		'ổ': c_ộ,
		'ở': c_ợ,
		'ủ': c_ụ,
		'ử': c_ự,
		'ỷ': c_ỵ,

		'ã': c_ạ,
		'ẵ': c_ặ,
		'ẫ': c_ậ,
		'ẽ': c_ẹ,
		'ễ': c_ệ,
		'ĩ': c_ị,
		'õ': c_ọ,
		'ỗ': c_ộ,
		'ỡ': c_ợ,
		'ũ': c_ụ,
		'ữ': c_ự,
		'ỹ': c_ỵ,

		'ạ': r_a,
		'ặ': r_ă,
		'ậ': r_â,
		'ẹ': r_e,
		'ệ': r_ê,
		'ị': r_i,
		'ọ': r_o,
		'ộ': r_ô,
		'ợ': r_ơ,
		'ụ': r_u,
		'ự': r_ư,
		'ỵ': r_y,
	},
	'd': {
		'd': c_đ,
		'đ': r_d,
	},
	'a': {
		'a': c_â,
		'á': c_ấ,
		'à': c_ầ,
		'ả': c_ẩ,
		'ã': c_ẫ,
		'ạ': c_ậ,

		'ă': c_â,
		'ắ': c_ấ,
		'ằ': c_ầ,
		'ẳ': c_ẩ,
		'ẵ': c_ẫ,
		'ặ': c_ậ,

		'â': r_a,
		'ấ': r_á,
		'ầ': r_à,
		'ẩ': r_ả,
		'ẫ': r_ã,
		'ậ': r_ạ,
	},
	'o': {
		'o': c_ô,
		'ó': c_ố,
		'ò': c_ồ,
		'ỏ': c_ổ,
		'õ': c_ỗ,
		'ọ': c_ộ,

		'ơ': c_ô,
		'ớ': c_ố,
		'ờ': c_ồ,
		'ở': c_ổ,
		'ỡ': c_ỗ,
		'ợ': c_ộ,

		'ô': r_o,
		'ố': r_ó,
		'ồ': r_ò,
		'ổ': r_ỏ,
		'ỗ': r_õ,
		'ộ': r_ọ,
	},
	'e': {
		'e': c_ê,
		'é': c_ế,
		'è': c_ề,
		'ẻ': c_ể,
		'ẽ': c_ễ,
		'ẹ': c_ệ,

		'ê': r_e,
		'ế': r_é,
		'ề': r_è,
		'ể': r_ẻ,
		'ễ': r_ẽ,
		'ệ': r_ẹ,
	},

	'w': {
		'a': c_ă,
		'á': c_ắ,
		'à': c_ằ,
		'ả': c_ẳ,
		'ã': c_ẵ,
		'ạ': c_ặ,

		'â': c_ă,
		'ấ': c_ắ,
		'ầ': c_ằ,
		'ẩ': c_ẳ,
		'ẫ': c_ẵ,
		'ậ': c_ặ,

		'o': c_ơ,
		'ó': c_ớ,
		'ò': c_ờ,
		'ỏ': c_ở,
		'õ': c_ỡ,
		'ọ': c_ợ,

		'ô': c_ơ,
		'ố': c_ớ,
		'ồ': c_ờ,
		'ổ': c_ở,
		'ỗ': c_ỡ,
		'ộ': c_ợ,

		'u': c_ư,
		'ú': c_ứ,
		'ù': c_ừ,
		'ủ': c_ử,
		'ũ': c_ữ,
		'ụ': c_ự,

		'ă': r_a,
		'ắ': r_á,
		'ằ': r_à,
		'ẳ': r_ả,
		'ẵ': r_ã,
		'ặ': r_ạ,

		'ơ': r_o,
		'ờ': r_ó,
		'ớ': r_ò,
		'ở': r_ỏ,
		'ỡ': r_õ,
		'ợ': r_ọ,

		'ư': r_u,
		'ứ': r_ú,
		'ừ': r_ù,
		'ử': r_ủ,
		'ữ': r_ũ,
		'ự': r_ụ,
	},
	'z': {
		'á': c_a,
		'ắ': c_ă,
		'ấ': c_â,
		'é': c_e,
		'ế': c_ê,
		'í': c_i,
		'ó': c_o,
		'ố': c_ô,
		'ớ': c_ơ,
		'ú': c_u,
		'ứ': c_ư,
		'ý': c_y,
		'à': c_a,
		'ằ': c_ă,
		'ầ': c_â,
		'è': c_e,
		'ề': c_ê,
		'ì': c_i,
		'ò': c_o,
		'ồ': c_ô,
		'ờ': c_ơ,
		'ù': c_u,
		'ừ': c_ư,
		'ỳ': c_y,
		'ả': c_a,
		'ẳ': c_ă,
		'ẩ': c_â,
		'ẻ': c_e,
		'ể': c_ê,
		'ỉ': c_i,
		'ỏ': c_o,
		'ổ': c_ô,
		'ở': c_ơ,
		'ủ': c_u,
		'ử': c_ư,
		'ỷ': c_y,
		'ã': c_a,
		'ẵ': c_ă,
		'ẫ': c_â,
		'ẽ': c_e,
		'ễ': c_ê,
		'ĩ': c_i,
		'õ': c_o,
		'ỗ': c_ô,
		'ỡ': c_ơ,
		'ũ': c_u,
		'ữ': c_ư,
		'ỹ': c_y,
		'ạ': c_a,
		'ặ': c_ă,
		'ậ': c_â,
		'ẹ': c_e,
		'ệ': c_ê,
		'ị': c_i,
		'ọ': c_o,
		'ộ': c_ô,
		'ợ': c_ơ,
		'ụ': c_u,
		'ự': c_ư,
		'ỵ': c_y,
	},
}
