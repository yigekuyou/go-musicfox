package ui

// IMenu menu interface
type IMenu interface {
    // IsPlayable 当前菜单是否可播放？
    IsPlayable() bool

    // ResetPlaylistWhenEnter 进入当前菜单前，是否重置播放列表？
    ResetPlaylistWhenEnter() bool

    // GetMenuKey 菜单唯一Key
    GetMenuKey() string

    // GetSubMenuViews 获取子菜单View
    GetSubMenuViews() []string

    // SubMenu 根据下标获取菜单Model
    SubMenu(index int) IMenu

    // ExtraView 获取额外的View（只在getSubMenuViews返回空时才会渲染到菜单的位置）
    ExtraView() string

    // BeforePrePageHook 切换上一页前的Hook
    BeforePrePageHook(m *NeteaseModel)

    // BeforeNextPageHook 切换下一页前的Hook
    BeforeNextPageHook(m *NeteaseModel)

    // BeforeEnterMenuHook 进入菜单项前的Hook
    BeforeEnterMenuHook(m *NeteaseModel) []string

    // BottomOutHook 触底的Hook
    BottomOutHook(m *NeteaseModel)

    // TopOutHook 触顶Hook
    TopOutHook(m *NeteaseModel)
}

// IDjMenu dj menu interface
type IDjMenu interface {
    IMenu
}