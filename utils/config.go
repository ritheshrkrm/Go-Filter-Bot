// (c) Jisin0

package utils

import "github.com/PaulSonOfLars/gotgbot/v2"

var TEXT map[string]string = map[string]string{

	"START": `𝖧𝖾𝗅𝗅𝗈 𝖳𝗁𝖾𝗋𝖾,
𝖨'𝗆 𝗃𝗎𝗌𝗍 𝖺 𝗌𝗂𝗆𝗉𝗅𝖾 𝗆𝖺𝗇𝗎𝖺𝗅 𝖿𝗂𝗅𝗍𝖾𝗋 𝖻𝗈𝗍

𝗂𝗍𝗌 𝖾𝖺𝗌𝗒 𝗍𝗈 𝗎𝗌𝖾 𝗆𝖾; 𝗃𝗎𝗌𝗍 𝖺𝖽𝖽 𝗆𝖾 𝗍𝗈 𝗒𝗈𝗎𝗋 𝗀𝗋𝗈𝗎𝗉 𝖺𝗌 𝖺𝖽𝗆𝗂𝗇 𝗁𝗂𝗍 /help 𝖿𝗈𝗋 𝗆𝗈𝗋𝖾.
	`,
	"ABOUT": `○ 𝖢𝗋𝖾𝖺𝗍𝗈𝗋 : <a href='https://t.me/rithesh_rkrm_17'>𝖳𝗁𝗂𝗌 𝖯𝖾𝗋𝗌𝗈𝗇</a>
○ 𝖫𝖺𝗇𝗀𝗎𝖺𝗀𝖾 : 𝖯𝗒𝗍𝗁𝗈𝗇 𝟥 
○ 𝖫𝗂𝖻𝗋𝖺𝗋𝗒 : 𝖯𝗒𝗋𝗈𝗀𝗋𝖺𝗆 𝖺𝗌𝗒𝗇𝖼𝗂𝗈 𝟢.𝟣𝟩.𝟣 
○ 𝖲𝖾𝗋𝗏𝖾𝗋 : Contabo
○ 𝖣𝖺𝗍𝖺𝖻𝖺𝗌𝖾 : <a href='https://www.mongodb.com'>𝖬𝗈𝗇𝗀𝗈𝖣𝖡 𝖥𝗋𝖾𝖾 𝖳𝗂𝖾𝗋</a>
○ 𝖡𝗎𝗂𝗅𝖽 𝖲𝗍𝖺𝗍𝗎𝗌 : v1.0.1 [BeTa]
○ 𝖲𝗎𝗉𝗉𝗈𝗋𝗍 𝖦𝗋𝗈𝗎𝗉 : <a href='https://t.me/raixchat'>𝖳𝖺𝗉 𝖧𝖾𝗋𝖾</a>
	`,

	"MF": `
	Mᴀɴᴜᴀʟ ғɪʟᴛᴇʀs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ sᴀᴠᴇ ᴄᴜsᴛᴏᴍ ғɪʟᴛᴇʀs ᴏᴛʜᴇʀ ᴛʜᴀɴ ᴛʜᴇ ᴀᴜᴛᴏᴍᴀᴛɪᴄ ᴏɴᴇs. 🌪 𝖥𝗂𝗅𝗍𝖾𝗋s ᴄᴀɴ ʙᴇ ᴏғ ᴛᴇxᴛ/ᴘʜᴏᴛᴏ/ᴅᴏᴄᴜᴍᴇɴᴛ/ᴀᴜᴅɪᴏ/ᴀɴɪᴍᴀᴛɪᴏɴ/ᴠɪᴅᴇᴏ .   
   
Nᴇᴡ ғɪʟᴛᴇʀ :   
   
Fᴏʀᴍᴀᴛ   
/filter "keyword" text or   
Rᴇᴘʟʏ ᴛᴏ ᴀ ᴍᴇssᴀɢᴇ ->/filter "keyword"   
Usᴀɢᴇ   
/filter "hi" hello   
[hello] -> Reply -> /filter hi   
   
Sᴛᴏᴘ ғɪʟᴛᴇʀ :   
   
Fᴏʀᴍᴀᴛ   
/stop "keyword"   
Usᴀɢᴇ   
/stop "hi"   
   
Vɪᴇᴡ ғɪʟᴛᴇʀs :   
/filters
`,

	"GF": `
	Gʟᴏʙᴀʟ ғɪʟᴛᴇʀs ᴀʀᴇ ᴍᴀɴᴜᴀʟ ғɪʟᴛᴇʀs sᴀᴠᴇᴅ ʙʏ ʙᴏᴛ ᴀᴅᴍɪɴs ᴛʜᴀᴛ ᴡᴏʀᴋ ɪɴ ᴀʟʟ ᴄʜᴀᴛs. Tʜᴇʏ ᴘʀᴏᴠɪᴅᴇ ʟᴀᴛᴇsᴛ ᴍᴏᴠɪᴇs ɪɴ ᴀ ᴇᴀsʏ ᴛᴏ ᴜsᴇ ғᴏʀᴍᴀᴛ.  
  
	Sᴛᴏᴘ ғɪʟᴛᴇʀ :  
	  
	Fᴏʀᴍᴀᴛ  
	/stop "keyword"  
	Usᴀɢᴇ  
	/stop "et"  
	  
	Vɪᴇᴡ ғɪʟᴛᴇʀs :  
	/gfilters
`,
	"CONNECT": `
	Help: <b>Connections</b>

	- Used to connect bot to PM for managing filters 
	- it helps to avoid spamming in groups.
	
	<b>NOTE:</b>
	1. Only admins can add a connection.
	2. Send <code>/connect</code> for connecting me to ur PM
	
	<b>Commands and Usage:</b>
	• /connect  - <code>connect a particular chat to your PM</code>
	• /disconnect  - <code>disconnect from a chat</code>
	• /connections - <code>list all your connections</code>
`,

	"BROADCAST": `
	𝖳𝗁𝖾 𝖻𝗋𝗈𝖺𝖽𝖼𝖺𝗌𝗍 𝖿𝖾𝖺𝗍𝗎𝗋𝖾 𝖺𝗅𝗅𝗈𝗐𝗌 𝖻𝗈𝗍 𝖺𝖽𝗆𝗂𝗇𝗌 𝗍𝗈 𝖻𝗋𝗈𝖺𝖽𝖼𝖺𝗌𝗍 𝖺 𝗆𝖾𝗌𝗌𝖺𝗀𝖾 𝗍𝗈 𝖺𝗅𝗅 𝗈𝖿 𝗍𝗁𝖾 𝖻𝗈𝗍'𝗌 𝗎𝗌𝖾𝗋𝗌.  
  
	𝖡𝗋𝗈𝖺𝖽𝖼𝖺𝗌𝗍𝗌 𝖺𝗋𝖾 𝗈𝖿 𝗍𝗐𝗈 𝗍𝗒𝗉𝖾𝗌 :  
	 ~ 𝖥𝗎𝗅𝗅 𝖡𝗋𝗈𝖺𝖽𝖼𝖺𝗌𝗍 - 𝖡𝗋𝗈𝖺𝖽𝖼𝖺𝗌𝗍 𝗍𝗈 𝖺𝗅𝗅 𝗈𝖿 𝗍𝗁𝖾 𝖻𝗈𝗍 𝗎𝗌𝖾𝗋𝗌 <code>/𝖻𝗋𝗈𝖺𝖽𝖼𝖺𝗌𝗍</code>  
	 ~ 𝖢𝗈𝗇𝖼𝖺𝗌𝗍 - 𝖡𝗋𝗈𝖺𝖽𝖼𝖺𝗌𝗍 𝗍𝗈 𝗈𝗇𝗅𝗒 𝗎𝗌𝖾𝗋𝗌 𝗐𝗁𝗈 𝖺𝗋𝖾 𝖼𝗈𝗇𝗇𝖾𝖼𝗍𝖾𝖽 𝗍𝗈 𝖺 𝖼𝗁𝖺𝗍 <code>/𝖼𝗈𝗇𝖼𝖺𝗌𝗍</code>
`,

	"HELP": `
	𝖳𝗈 𝗄𝗇𝗈𝗐 𝗁𝗈𝗐 𝗍𝗈 𝗎𝗌𝖾 𝗆𝗒 𝗆𝗈𝖽𝗎𝗅𝖾𝗌 𝗎𝗌𝖾 𝗍𝗁𝖾 𝖻𝗎𝗍𝗍𝗈𝗇𝗌 𝖻𝖾𝗅𝗈𝗐 𝗍𝗈 𝗀𝖾𝗍 𝖺𝗅𝗅 𝗆𝗒 𝖼𝗈𝗆𝗆𝖺𝗇𝖽𝗌 𝗐𝗂𝗍𝗁 𝗎𝗌𝖺𝗀𝖾 𝖾𝗑𝖺𝗆𝗉𝗅𝖾𝗌 !
`,
}

var BUTTONS map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"START": {
		{
			{Text: "😊 𝖠𝖻𝗈𝗎𝗍", CallbackData: "edit(ABOUT)"},
			{Text: "ℹ️ 𝖧𝖾𝗅𝗉", CallbackData: "edit(HELP)"},
			{Text: "🌐 𝖲𝗎𝗉𝗉𝗈𝗋𝗍 ", Url: "t.me/raixchat"},
		},
	},
	"ABOUT": {
		{
			{Text: "🏘 Home", CallbackData: "edit(START)"},
			{Text: "🧩 Stats", CallbackData: "stats"},
		},
	},
	"STATS": {
		{
			{Text: "◀️ 𝖡𝖠𝖢𝖪", CallbackData: "edit(ABOUT)"},
			{Text: "🧩 Stats", CallbackData: "stats"},
		},
	},
	"HELP": {
		{{Text: "🌪 𝖥𝗂𝗅𝗍𝖾𝗋", CallbackData: "edit(MF)"},
			{Text: "🌍 𝖦𝗅𝗈𝖻𝖺𝗅", CallbackData: "edit(GF)"},
		}, {
			{Text: "🖥 𝖢𝗈𝗇𝗇𝖾𝖼𝗍", CallbackData: "edit(CONNECT)"}, {Text: "Broadcast", CallbackData: "edit(BROADCAST)"},
		},
		{{Text: "◀️ 𝖡𝖠𝖢𝖪", CallbackData: "edit(START)"}},
	},
}
