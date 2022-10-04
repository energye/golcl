package locales

import "github.com/energye/golcl/lcl/api"

// 修改资源
func ModifyResources(data map[string]string) {
	for i := int32(0); i < api.DGetLibResourceCount(); i++ {
		item := api.DGetLibResourceItem(i)
		if value, ok := data[item.Name]; ok {
			api.DModifyLibResource(item.Ptr, value)
		}
	}
}
