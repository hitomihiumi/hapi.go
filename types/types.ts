interface Responce<T> {
    data: T;
    status: number;
}

interface DiscordUser {
    id: string;
    email: string;
    global_name: string;
    username: string;
    discriminator: string;
    avatar: string;
    avatar_url: string;
    banner: string;
    banner_url: string;
    locale: string;
    mfa_enabled: boolean;
    verified: boolean;
    accent_color: number;
    bot: boolean;
    public_flags: number;
    premium_type: number;
    system: boolean;
    flags: number;
}

interface DiscordActivity {
    name: string;
    type: number;
    created_at: string;
    application_id?: string;
    details?: string;
    state?: string;
    assets?: {
        large_image?: string;
        large_text?: string;
        small_image?: string;
        small_text?: string;
    };
    timestamps: {
        start?: string;
        end?: string;
    };
    emoji: {
        id: string;
        name: string;
        roles: null | string[];
        user: null | DiscordUser;
        require_colons: boolean;
        managed: boolean;
        animated: boolean;
        available: boolean;
    };
    party: {
        id?: string;
        size?: number[];
    };
}

interface SteamUser {
    steamid: string;
    communityvisibilitystate: number;
    profilestate: number;
    personaname: string;
    lastlogoff: number;
    commentpermission: number;
    profileurl: string;
    avatar: string;
    avatarmedium: string;
    avatarfull: string;
    avatarhash: string;
    personastate: number;
    realname: string;
    primaryclanid?: string;
    timecreated: number;
    personastateflags: number;
    loccountrycode: string;
    locstatecode: string;
    loccityid: number;
    background: string;
    frame: string;
    level: number;
}

interface SteamGameRecentlyPlayed {
    appid: number;
    name: string;
    playtime_2weeks: number;
    playtime_forever: number;
    img_icon_url: string;
    img_logo_url: string;
    playtime_windows_forever: number;
    playtime_mac_forever: number;
    playtime_linux_forever: number;
    assets: GameAssets;
}

interface SteamGameOwned {
    appid: number;
    name: string;
    playtime_forever: number;
    playtime_windows_forever: number;
    playtime_mac_forever: number;
    playtime_linux_forever: number;
    img_icon_url: string;
    img_logo_url: string;
    has_community_visible_stats: boolean;
    assets: GameAssets;
}

interface GameAssets {
    capsule: string;
    capsule_2x: string;
    small_capsule: string;
    small_capsule_2x: string
    header: string;
    header_2x: string;
    logo: string;
    logo_2x: string;
    hero: string;
    hero_2x: string;
}

interface SteamGameAchievements {
    steamID: string;
    gameName: string;
    achievements: {
        apiname: string;
        achieved: number;
        unlocktime: number;
        name?: string;
        description?: string;
    }[];
}

export {
    Responce,
    DiscordUser,
    DiscordActivity,
    SteamUser,
    SteamGameRecentlyPlayed,
    SteamGameOwned,
    SteamGameAchievements,
    GameAssets
}