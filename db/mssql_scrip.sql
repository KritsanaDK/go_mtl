USE [ksn_db]
GO
/****** Object:  Table [dbo].[itemData]    Script Date: 2025-01-29 23:20:22 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[itemData](
	[id] [int] IDENTITY(0,1) NOT NULL,
	[item] [varchar](100) NULL
) ON [PRIMARY]
GO
SET IDENTITY_INSERT [dbo].[itemData] ON 

INSERT [dbo].[itemData] ([id], [item]) VALUES (26, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (9, N'Updated Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (10, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (11, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (12, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (13, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (14, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (15, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (16, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (17, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (18, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (19, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (20, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (21, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (22, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (23, N'Sample Item')
INSERT [dbo].[itemData] ([id], [item]) VALUES (24, N'Sample Item')
SET IDENTITY_INSERT [dbo].[itemData] OFF
GO
