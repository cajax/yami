package yami

type TrackType string

const (
	TrackGeneral TrackType = "General"
	TrackVideo   TrackType = "Video"
	TrackAudio   TrackType = "Audio"
)

type MediaInfo struct {
	Media *Media `xml:"media" json:"media"`
}

type Media struct {
	Tracks []*Track `xml:"track" json:"tracks"`
}

type Track struct {
	Album_Sort                          string    `xml:"Album_Sort" json:"album_sort,omitempty"`
	Alignment                           string    `xml:"Alignment" json:"alignment"`
	AlternateGroup                      string    `xml:"AlternateGroup" json:"alternate_group,omitempty"`
	Archival_Location                   string    `xml:"Archival_Location" json:"archival_location,omitempty"`
	Arranger                            string    `xml:"Arranger" json:"arranger"`
	ArtDirector                         string    `xml:"ArtDirector" json:"art_director"`
	AssistantDirector                   string    `xml:"AssistantDirector" json:"assistant_director,omitempty"`
	AudioCount                          int       `xml:"AudioCount" json:"audio_count,omitempty"`
	BarCode                             string    `xml:"BarCode" json:"bar_code,omitempty"`
	BitDepth                            int       `xml:"BitDepth" json:"bit_depth,omitempty"`
	BitDepth_Detected                   int       `xml:"BitDepth_Detected" json:"bit_depth_detected,omitempty"`
	BitDepth_Stored                     int       `xml:"BitDepth_Stored" json:"bit_depth_stored,omitempty"`
	BitRate                             float64   `xml:"BitRate" json:"bit_rate,omitempty"`
	BitRate_Encoded                     float64   `xml:"BitRate_Encoded" json:"bit_rate_encoded,omitempty"`
	BitRate_Maximum                     float64   `xml:"BitRate_Maximum" json:"bit_rate_maximum,omitempty"`
	BitRate_Minimum                     float64   `xml:"BitRate_Minimum" json:"bit_rate_minimum,omitempty"`
	BitRate_Mode                        string    `xml:"BitRate_Mode" json:"bit_rate_mode,omitempty"`
	BitRate_Nominal                     float64   `xml:"BitRate_Nominal" json:"bit_rate_nominal,omitempty"`
	BPM                                 string    `xml:"BPM" json:"bPM,omitempty"`
	BufferSize                          string    `xml:"BufferSize" json:"buffer_size,omitempty"`
	CatalogNumber                       string    `xml:"CatalogNumber" json:"catalog_number,omitempty"`
	ChannelLayout                       string    `xml:"ChannelLayout" json:"channel_layout,omitempty"`
	ChannelPositions                    string    `xml:"ChannelPositions" json:"channel_positions,omitempty"`
	Channels                            int       `xml:"Channels" json:"channels,omitempty"`
	Channels_Original                   int       `xml:"Channels_Original" json:"channels_original,omitempty"`
	Chapter                             string    `xml:"Chapter" json:"chapter,omitempty"`
	Choregrapher                        string    `xml:"Choregrapher" json:"choregrapher,omitempty"`
	ChromaSubsampling                   string    `xml:"ChromaSubsampling" json:"chroma_subsampling,omitempty"`
	ChromaSubsampling_Position          string    `xml:"ChromaSubsampling_Position" json:"chroma_subsampling_position,omitempty"`
	CodecID                             string    `xml:"CodecID" json:"codec_id,omitempty"`
	CodecID_Description                 string    `xml:"CodecID_Description" json:"codec_id_description,omitempty"`
	CodecID_Version                     string    `xml:"CodecID_Version" json:"codec_id_version,omitempty"`
	CodecID_Compatible                  string    `xml:"CodecID_Compatible" json:"codec_id_compatible,omitempty"`
	CoDirector                          string    `xml:"CoDirector" json:"co_director,omitempty"`
	Collection                          string    `xml:"Collection" json:"collection,omitempty"`
	ColorSpace                          string    `xml:"ColorSpace" json:"color_space,omitempty"`
	Colour_description_present          string    `xml:"colour_description_present" json:"colour_description_present,omitempty"`
	Colour_description_present_Original string    `xml:"colour_description_present_Original" json:"colour_description_present_original,omitempty"`
	Colour_description_present_Source   string    `xml:"colour_description_present_Source" json:"colour_description_present_source,omitempty"`
	Colour_primaries                    string    `xml:"colour_primaries" json:"colour_primaries,omitempty"`
	Colour_primaries_Original           string    `xml:"colour_primaries_Original" json:"colour_primaries_original,omitempty"`
	Colour_primaries_Source             string    `xml:"colour_primaries_Source" json:"colour_primaries_source,omitempty"`
	Colour_range                        string    `xml:"colour_range" json:"colour_range,omitempty"`
	Colour_range_Source                 string    `xml:"colour_range_Source" json:"colour_range_source,omitempty"`
	Comic                               string    `xml:"Comic" json:"comic,omitempty"`
	Comic_More                          string    `xml:"Comic_More" json:"comic_more,omitempty"`
	Comic_Position_Total                int       `xml:"Comic_Position_Total" json:"comic_position_total,omitempty"`
	Comment                             string    `xml:"Comment" json:"comment,omitempty"`
	CommissionedBy                      string    `xml:"CommissionedBy" json:"commissioned_by,omitempty"`
	Compilation                         string    `xml:"Compilation" json:"compilation,omitempty"`
	CompleteName                        string    `xml:"CompleteName" json:"complete_name,omitempty"`
	CompleteName_Last                   string    `xml:"CompleteName_Last" json:"complete_name_last,omitempty"`
	Composer                            string    `xml:"Composer" json:"composer,omitempty"`
	Composer_Nationality                string    `xml:"Composer_Nationality" json:"composer_nationality,omitempty"`
	Composer_Sort                       string    `xml:"Composer_Sort" json:"composer_sort,omitempty"`
	Compression_Mode                    string    `xml:"Compression_Mode" json:"compression_mode,omitempty"`
	Compression_Ratio                   float64   `xml:"Compression_Ratio" json:"compression_ratio,omitempty"`
	Conductor                           string    `xml:"Conductor" json:"conductor,omitempty"`
	ContentType                         string    `xml:"ContentType" json:"content_type,omitempty"`
	CoProducer                          string    `xml:"CoProducer" json:"co_producer,omitempty"`
	Copyright                           string    `xml:"Copyright" json:"copyright,omitempty"`
	Copyright_Url                       string    `xml:"Copyright_Url" json:"copyright_url,omitempty"`
	CostumeDesigner                     string    `xml:"CostumeDesigner" json:"costume_designer,omitempty"`
	Countries                           string    `xml:"Countries" json:"countries,omitempty"`
	Country                             string    `xml:"Country" json:"country,omitempty"`
	Cover                               string    `xml:"Cover" json:"cover,omitempty"`
	Cover_Data                          string    `xml:"Cover_Data" json:"cover_data,omitempty"`
	Cover_Description                   string    `xml:"Cover_Description" json:"cover_description,omitempty"`
	Cover_Mime                          string    `xml:"Cover_Mime" json:"cover_mime,omitempty"`
	Cover_Type                          string    `xml:"Cover_Type" json:"cover_type,omitempty"`
	Cropped                             string    `xml:"Cropped" json:"cropped,omitempty"`
	DataSize                            int64       `xml:"DataSize" json:"data_size,omitempty"`
	Default                             string    `xml:"Default" json:"default,omitempty"`
	Delay                               float64   `xml:"Delay" json:"delay,omitempty"`
	Delay_DropFrame                     string    `xml:"Delay_DropFrame" json:"delay_drop_frame,omitempty"`
	Delay_Original                      float64   `xml:"Delay_Original" json:"delay_original,omitempty"`
	Delay_Original_DropFrame            string    `xml:"Delay_Original_DropFrame" json:"delay_original_drop_frame,omitempty"`
	Delay_Original_Source               string    `xml:"Delay_Original_Source" json:"delay_original_source,omitempty"`
	Delay_Settings                      string    `xml:"Delay_Settings" json:"delay_settings,omitempty"`
	Delay_Source                        string    `xml:"Delay_Source" json:"delay_source,omitempty"`
	Description                         string    `xml:"Description" json:"description,omitempty"`
	Dimensions                          string    `xml:"Dimensions" json:"dimensions,omitempty"`
	Director                            string    `xml:"Director" json:"director,omitempty"`
	DirectorOfPhotography               string    `xml:"DirectorOfPhotography" json:"director_of_photography,omitempty"`
	Disabled                            string    `xml:"Disabled" json:"disabled,omitempty"`
	DisplayAspectRatio                  float64   `xml:"DisplayAspectRatio" json:"display_aspect_ratio,omitempty"`
	DisplayAspectRatio_Original         float64   `xml:"DisplayAspectRatio_Original" json:"display_aspect_ratio_original,omitempty"`
	DistributedBy                       string    `xml:"DistributedBy" json:"distributed_by,omitempty"`
	Domain                              string    `xml:"Domain" json:"domain,omitempty"`
	DotsPerInch                         string    `xml:"DotsPerInch" json:"dots_per_inch,omitempty"`
	Duration                            float64   `xml:"Duration" json:"duration,omitempty"`
	Duration_End                        float64   `xml:"Duration_End" json:"duration_end,omitempty"`
	Duration_FirstFrame                 float64   `xml:"Duration_FirstFrame" json:"duration_first_frame,omitempty"`
	Duration_LastFrame                  float64   `xml:"Duration_LastFrame" json:"duration_last_frame,omitempty"`
	Duration_Start                      float64   `xml:"Duration_Start" json:"duration_start,omitempty"`
	EditedBy                            string    `xml:"EditedBy" json:"edited_by,omitempty"`
	ElementCount                        int       `xml:"ElementCount" json:"element_count,omitempty"`
	Encoded_Application                 string    `xml:"Encoded_Application" json:"encoded_application,omitempty"`
	Encoded_Application_CompanyName     string    `xml:"Encoded_Application_CompanyName" json:"encoded_application_company_name,omitempty"`
	Encoded_Application_Name            string    `xml:"Encoded_Application_Name" json:"encoded_application_name,omitempty"`
	Encoded_Application_String          string    `xml:"Encoded_Application_String" json:"encoded_application_string,omitempty"`
	Encoded_Application_Url             string    `xml:"Encoded_Application_Url" json:"encoded_application_url,omitempty"`
	Encoded_Application_Version         string    `xml:"Encoded_Application_Version" json:"encoded_application_version,omitempty"`
	Encoded_Date                        string    `xml:"Encoded_Date" json:"encoded_date,omitempty"`
	Encoded_Library                     string    `xml:"Encoded_Library" json:"encoded_library,omitempty"`
	Encoded_Library_CompanyName         string    `xml:"Encoded_Library_CompanyName" json:"encoded_library_company_name,omitempty"`
	Encoded_Library_Date                string    `xml:"Encoded_Library_Date" json:"encoded_library_date,omitempty"`
	Encoded_Library_Name                string    `xml:"Encoded_Library_Name" json:"encoded_library_name,omitempty"`
	Encoded_Library_Settings            string    `xml:"Encoded_Library_Settings" json:"encoded_library_settings,omitempty"`
	Encoded_Library_Version             string    `xml:"Encoded_Library_Version" json:"encoded_library_version,omitempty"`
	Encoded_OperatingSystem             string    `xml:"Encoded_OperatingSystem" json:"encoded_operating_system,omitempty"`
	EncodedBy                           string    `xml:"EncodedBy" json:"encoded_by,omitempty"`
	Encryption                          string    `xml:"Encryption" json:"encryption,omitempty"`
	Encryption_Format                   string    `xml:"Encryption_Format" json:"encryption_format,omitempty"`
	Encryption_InitializationVector     string    `xml:"Encryption_InitializationVector" json:"encryption_initialization_vector,omitempty"`
	Encryption_Length                   string    `xml:"Encryption_Length" json:"encryption_length,omitempty"`
	Encryption_Method                   string    `xml:"Encryption_Method" json:"encryption_method,omitempty"`
	Encryption_Mode                     string    `xml:"Encryption_Mode" json:"encryption_mode,omitempty"`
	Encryption_Padding                  string    `xml:"Encryption_Padding" json:"encryption_padding,omitempty"`
	EPG_Positions_Begin                 int64       `xml:"EPG_Positions_Begin" json:"ePG_positions_begin,omitempty"`
	EPG_Positions_End                   int64      `xml:"EPG_Positions_End" json:"ePG_positions_end,omitempty"`
	ExecutiveProducer                   string    `xml:"ExecutiveProducer" json:"executive_producer,omitempty"`
	File_Created_Date                   string    `xml:"File_Created_Date" json:"file_created_date,omitempty"`
	File_Created_Date_Local             string    `xml:"File_Created_Date_Local" json:"file_created_date_local,omitempty"`
	File_Modified_Date                  string    `xml:"File_Modified_Date" json:"file_modified_date,omitempty"`
	File_Modified_Date_Local            string    `xml:"File_Modified_Date_Local" json:"file_modified_date_local,omitempty"`
	FileExtension                       string    `xml:"FileExtension" json:"file_extension,omitempty"`
	FileSize                            string    `xml:"FileSize" json:"file_size,omitempty"`
	FirstPacketOrder                    int       `xml:"FirstPacketOrder" json:"first_packet_order,omitempty"`
	FooterSize                          int64       `xml:"FooterSize" json:"footer_size,omitempty"`
	Forced                              string    `xml:"Forced" json:"forced,omitempty"`
	Format                              string    `xml:"Format" json:"format,omitempty"`
	Format_AdditionalFeatures           string    `xml:"Format_AdditionalFeatures" json:"format_additional_features,omitempty"`
	Format_Commercial_IfAny             string    `xml:"Format_Commercial_IfAny" json:"format_commercial_if_any,omitempty"`
	Format_Compression                  string    `xml:"Format_Compression" json:"format_compression,omitempty"`
	Format_Level                        string    `xml:"Format_Level" json:"format_level,omitempty"`
	Format_Profile                      string    `xml:"Format_Profile" json:"format_profile,omitempty"`
	Format_Settings                     string    `xml:"Format_Settings" json:"format_settings,omitempty"`
	Format_Settings_BVOP                string    `xml:"Format_Settings_BVOP" json:"format_settings_bVOP,omitempty"`
	Format_Settings_CABAC               string    `xml:"Format_Settings_CABAC" json:"format_settings_cABAC,omitempty"`
	Format_Settings_Emphasis            string    `xml:"Format_Settings_Emphasis" json:"format_settings_emphasis,omitempty"`
	Format_Settings_Endianness          string    `xml:"Format_Settings_Endianness" json:"format_settings_endianness,omitempty"`
	Format_Settings_Firm                string    `xml:"Format_Settings_Firm" json:"format_settings_firm,omitempty"`
	Format_Settings_Floor               string    `xml:"Format_Settings_Floor" json:"format_settings_floor,omitempty"`
	Format_Settings_FrameMode           string    `xml:"Format_Settings_FrameMode" json:"format_settings_frame_mode,omitempty"`
	Format_Settings_GMC                 int       `xml:"Format_Settings_GMC" json:"format_settings_gMC,omitempty"`
	Format_Settings_GOP                 string    `xml:"Format_Settings_GOP" json:"format_settings_gOP,omitempty"`
	Format_Settings_ITU                 string    `xml:"Format_Settings_ITU" json:"format_settings_iTU,omitempty"`
	Format_Settings_Law                 string    `xml:"Format_Settings_Law" json:"format_settings_law,omitempty"`
	Format_Settings_Matrix              string    `xml:"Format_Settings_Matrix" json:"format_settings_matrix,omitempty"`
	Format_Settings_Matrix_Data         string    `xml:"Format_Settings_Matrix_Data" json:"format_settings_matrix_data,omitempty"`
	Format_Settings_Mode                string    `xml:"Format_Settings_Mode" json:"format_settings_mode,omitempty"`
	Format_Settings_ModeExtension       string    `xml:"Format_Settings_ModeExtension" json:"format_settings_mode_extension,omitempty"`
	Format_Settings_Packing             string    `xml:"Format_Settings_Packing" json:"format_settings_packing,omitempty"`
	Format_Settings_PictureStructure    string    `xml:"Format_Settings_PictureStructure" json:"format_settings_picture_structure,omitempty"`
	Format_Settings_PS                  string    `xml:"Format_Settings_PS" json:"format_settings_pS,omitempty"`
	Format_Settings_Pulldown            string    `xml:"Format_Settings_Pulldown" json:"format_settings_pulldown,omitempty"`
	Format_Settings_QPel                string    `xml:"Format_Settings_QPel" json:"format_settings_qPel,omitempty"`
	Format_Settings_RefFrames           int       `xml:"Format_Settings_RefFrames" json:"format_settings_ref_frames,omitempty"`
	Format_Settings_SBR                 string    `xml:"Format_Settings_SBR" json:"format_settings_sBR,omitempty"`
	Format_Settings_Sign                string    `xml:"Format_Settings_Sign" json:"format_settings_sign,omitempty"`
	Format_Settings_Wrapping            string    `xml:"Format_Settings_Wrapping" json:"format_settings_wrapping,omitempty"`
	Format_Tier                         string    `xml:"Format_Tier" json:"format_tier,omitempty"`
	Format_Version                      string    `xml:"Format_Version" json:"format_version,omitempty"`
	FrameCount                          int64       `xml:"FrameCount" json:"frame_count,omitempty"`
	FrameRate                           float64   `xml:"FrameRate" json:"frame_rate,omitempty"`
	FrameRate_Maximum                   float64   `xml:"FrameRate_Maximum" json:"frame_rate_maximum,omitempty"`
	FrameRate_Minimum                   float64   `xml:"FrameRate_Minimum" json:"frame_rate_minimum,omitempty"`
	FrameRate_Mode                      string    `xml:"FrameRate_Mode" json:"frame_rate_mode,omitempty"`
	FrameRate_Mode_Original             string    `xml:"FrameRate_Mode_Original" json:"frame_rate_mode_original,omitempty"`
	FrameRate_Nominal                   float64   `xml:"FrameRate_Nominal" json:"frame_rate_nominal,omitempty"`
	FrameRate_Original                  float64   `xml:"FrameRate_Original" json:"frame_rate_original,omitempty"`
	GeneralCount                        int       `xml:"GeneralCount" json:"general_count,omitempty"`
	Genre                               string    `xml:"Genre" json:"genre,omitempty"`
	Gop_OpenClosed                      string    `xml:"Gop_OpenClosed" json:"gop_open_closed,omitempty"`
	Gop_OpenClosed_FirstFrame           string    `xml:"Gop_OpenClosed_FirstFrame" json:"gop_open_closed_first_frame,omitempty"`
	Grouping                            string    `xml:"Grouping" json:"grouping,omitempty"`
	HeaderSize                          int64       `xml:"HeaderSize" json:"header_size,omitempty"`
	Height                              int       `xml:"Height" json:"height,omitempty"`
	Height_Offset                       int       `xml:"Height_Offset" json:"height_offset,omitempty"`
	Height_Original                     int       `xml:"Height_Original" json:"height_original,omitempty"`
	ICRA                                string    `xml:"ICRA" json:"iCRA,omitempty"`
	ID                                  string    `xml:"ID" json:"iD,omitempty"`
	ImageCount                          int       `xml:"ImageCount" json:"image_count,omitempty"`
	Interleave_Duration                 float64   `xml:"Interleave_Duration" json:"interleave_duration,omitempty"`
	Interleave_Preload                  float64   `xml:"Interleave_Preload" json:"interleave_preload,omitempty"`
	Interleave_VideoFrames              float64   `xml:"Interleave_VideoFrames" json:"interleave_video_frames,omitempty"`
	Interleaved                         string    `xml:"Interleaved" json:"interleaved,omitempty"`
	ISBN                                string    `xml:"ISBN" json:"iSBN,omitempty"`
	ISRC                                string    `xml:"ISRC" json:"iSRC,omitempty"`
	IsStreamable                        string    `xml:"IsStreamable" json:"is_streamable,omitempty"`
	Keywords                            string    `xml:"Keywords" json:"keywords,omitempty"`
	Label                               string    `xml:"Label" json:"label,omitempty"`
	LabelCode                           string    `xml:"LabelCode" json:"label_code,omitempty"`
	Language                            string    `xml:"Language" json:"language,omitempty"`
	Language_More                       string    `xml:"Language_More" json:"language_more,omitempty"`
	LawRating                           string    `xml:"LawRating" json:"law_rating,omitempty"`
	LawRating_Reason                    string    `xml:"LawRating_Reason" json:"law_rating_reason,omitempty"`
	LCCN                                string    `xml:"LCCN" json:"lCCN,omitempty"`
	Lightness                           string    `xml:"Lightness" json:"lightness,omitempty"`
	List_StreamKind                     string    `xml:"List_StreamKind" json:"list_stream_kind,omitempty"`
	List_StreamPos                      string    `xml:"List_StreamPos" json:"list_stream_pos,omitempty"`
	Lyricist                            string    `xml:"Lyricist" json:"lyricist,omitempty"`
	Lyrics                              string    `xml:"Lyrics" json:"lyrics,omitempty"`
	Mastered_Date                       string    `xml:"Mastered_Date" json:"mastered_date,omitempty"`
	MasteredBy                          string    `xml:"MasteredBy" json:"mastered_by,omitempty"`
	Matrix_ChannelPositions             string    `xml:"Matrix_ChannelPositions" json:"matrix_channel_positions,omitempty"`
	Matrix_Channels                     int       `xml:"Matrix_Channels" json:"matrix_channels,omitempty"`
	Matrix_coefficients                 string    `xml:"matrix_coefficients" json:"matrix_coefficients,omitempty"`
	Matrix_coefficients_Original        string    `xml:"matrix_coefficients_Original" json:"matrix_coefficients_original,omitempty"`
	Matrix_coefficients_Source          string    `xml:"matrix_coefficients_Source" json:"matrix_coefficients_source,omitempty"`
	Matrix_Format                       string    `xml:"Matrix_Format" json:"matrix_format,omitempty"`
	MenuCount                           int       `xml:"MenuCount" json:"menu_count,omitempty"`
	MenuID                              string    `xml:"MenuID" json:"menu_id,omitempty"`
	Mood                                string    `xml:"Mood" json:"mood,omitempty"`
	Movie                               string    `xml:"Movie" json:"movie,omitempty"`
	Movie_Country                       string    `xml:"Movie_Country" json:"movie_country,omitempty"`
	Movie_More                          string    `xml:"Movie_More" json:"movie_more,omitempty"`
	MultiView_BaseProfile               string    `xml:"MultiView_BaseProfile" json:"multi_view_base_profile,omitempty"`
	MultiView_Count                     string    `xml:"MultiView_Count" json:"multi_view_count,omitempty"`
	MultiView_Layout                    string    `xml:"MultiView_Layout" json:"multi_view_layout,omitempty"`
	DolbyVision_Version                 string    `xml:"DolbyVision_Version" json:"dolby_vision_version,omitempty"`
	DolbyVision_Profile                 string    `xml:"DolbyVision_Profile" json:"dolby_vision_profile,omitempty"`
	DolbyVision_Layers                  string    `xml:"DolbyVision_Layers" json:"dolby_vision_layers,omitempty"`
	MusicBy                             string    `xml:"MusicBy" json:"music_by,omitempty"`
	MuxingMode                          string    `xml:"MuxingMode" json:"muxing_mode,omitempty"`
	MuxingMode_MoreInfo                 string    `xml:"MuxingMode_MoreInfo" json:"muxing_mode_more_info,omitempty"`
	NetworkName                         string    `xml:"NetworkName" json:"network_name,omitempty"`
	Original_Album                      string    `xml:"Original_Album" json:"original_album,omitempty"`
	Original_Lyricist                   string    `xml:"Original_Lyricist" json:"original_lyricist,omitempty"`
	Original_Movie                      string    `xml:"Original_Movie" json:"original_movie,omitempty"`
	Original_NetworkName                string    `xml:"Original_NetworkName" json:"original_network_name,omitempty"`
	Original_Part                       string    `xml:"Original_Part" json:"original_part,omitempty"`
	Original_Performer                  string    `xml:"Original_Performer" json:"original_performer,omitempty"`
	Original_Released_Date              string    `xml:"Original_Released_Date" json:"original_released_date,omitempty"`
	Original_Track                      string    `xml:"Original_Track" json:"original_track,omitempty"`
	OriginalNetworkName                 string    `xml:"OriginalNetworkName" json:"original_network_name,omitempty"`
	OriginalSourceForm                  string    `xml:"OriginalSourceForm" json:"original_source_form,omitempty"`
	OriginalSourceForm_Cropped          string    `xml:"OriginalSourceForm_Cropped" json:"original_source_form_cropped,omitempty"`
	OriginalSourceForm_DistributedBy    string    `xml:"OriginalSourceForm_DistributedBy" json:"original_source_form_distributed_by,omitempty"`
	OriginalSourceForm_Name             string    `xml:"OriginalSourceForm_Name" json:"original_source_form_name,omitempty"`
	OriginalSourceForm_NumColors        string    `xml:"OriginalSourceForm_NumColors" json:"original_source_form_num_colors,omitempty"`
	OriginalSourceForm_Sharpness        string    `xml:"OriginalSourceForm_Sharpness" json:"original_source_form_sharpness,omitempty"`
	OriginalSourceMedium                string    `xml:"OriginalSourceMedium" json:"original_source_medium,omitempty"`
	OriginalSourceMedium_ID             string    `xml:"OriginalSourceMedium_ID" json:"original_source_medium_id,omitempty"`
	OtherCount                          int       `xml:"OtherCount" json:"other_count,omitempty"`
	OverallBitRate                      float64   `xml:"OverallBitRate" json:"overall_bit_rate,omitempty"`
	OverallBitRate_Maximum              float64   `xml:"OverallBitRate_Maximum" json:"overall_bit_rate_maximum,omitempty"`
	OverallBitRate_Minimum              float64   `xml:"OverallBitRate_Minimum" json:"overall_bit_rate_minimum,omitempty"`
	OverallBitRate_Mode                 string    `xml:"OverallBitRate_Mode" json:"overall_bit_rate_mode,omitempty"`
	OverallBitRate_Nominal              float64   `xml:"OverallBitRate_Nominal" json:"overall_bit_rate_nominal,omitempty"`
	Owner                               string    `xml:"Owner" json:"owner,omitempty"`
	PackageName                         string    `xml:"PackageName" json:"package_name,omitempty"`
	Part                                string    `xml:"Part" json:"part,omitempty"`
	Part_Position                       int64       `xml:"Part_Position" json:"part_position,omitempty"`
	Part_Position_Total                 int64       `xml:"Part_Position_Total" json:"part_position_total,omitempty"`
	Performer                           string    `xml:"Performer" json:"performer,omitempty"`
	Performer_Sort                      string    `xml:"Performer_Sort" json:"performer_sort,omitempty"`
	Period                              string    `xml:"Period" json:"period,omitempty"`
	PixelAspectRatio                    float64   `xml:"PixelAspectRatio" json:"pixel_aspect_ratio,omitempty"`
	PixelAspectRatio_Original           float64   `xml:"PixelAspectRatio_Original" json:"pixel_aspect_ratio_original,omitempty"`
	Played_Count                        int64       `xml:"Played_Count" json:"played_count,omitempty"`
	Played_First_Date                   string    `xml:"Played_First_Date" json:"played_first_date,omitempty"`
	Played_Last_Date                    string    `xml:"Played_Last_Date" json:"played_last_date,omitempty"`
	PodcastCategory                     string    `xml:"PodcastCategory" json:"podcast_category,omitempty"`
	Producer                            string    `xml:"Producer" json:"producer,omitempty"`
	Producer_Copyright                  string    `xml:"Producer_Copyright" json:"producer_copyright,omitempty"`
	ProductionDesigner                  string    `xml:"ProductionDesigner" json:"production_designer,omitempty"`
	ProductionStudio                    string    `xml:"ProductionStudio" json:"production_studio,omitempty"`
	Publisher                           string    `xml:"Publisher" json:"publisher,omitempty"`
	Publisher_URL                       string    `xml:"Publisher_URL" json:"publisher_uRL,omitempty"`
	Rating                              string    `xml:"Rating" json:"rating,omitempty"`
	Recorded_Date                       string    `xml:"Recorded_Date" json:"recorded_date,omitempty"`
	Recorded_Location                   string    `xml:"Recorded_Location" json:"recorded_location,omitempty"`
	Released_Date                       string    `xml:"Released_Date" json:"released_date,omitempty"`
	RemixedBy                           string    `xml:"RemixedBy" json:"remixed_by,omitempty"`
	ReplayGain_Gain                     string    `xml:"ReplayGain_Gain" json:"replay_gain_gain,omitempty"`
	ReplayGain_Peak                     string    `xml:"ReplayGain_Peak" json:"replay_gain_peak,omitempty"`
	Rotation                            string    `xml:"Rotation" json:"rotation,omitempty"`
	Sampled_Height                      int       `xml:"Sampled_Height" json:"sampled_height,omitempty"`
	Sampled_Width                       int       `xml:"Sampled_Width" json:"sampled_width,omitempty"`
	SamplesPerFrame                     float64   `xml:"SamplesPerFrame" json:"samples_per_frame,omitempty"`
	SamplingCount                       int       `xml:"SamplingCount" json:"sampling_count,omitempty"`
	SamplingRate                        float64   `xml:"SamplingRate" json:"sampling_rate,omitempty"`
	ScanOrder                           string    `xml:"ScanOrder" json:"scan_order,omitempty"`
	ScanOrder_Original                  string    `xml:"ScanOrder_Original" json:"scan_order_original,omitempty"`
	ScanOrder_Stored                    string    `xml:"ScanOrder_Stored" json:"scan_order_stored,omitempty"`
	ScanType                            string    `xml:"ScanType" json:"scan_type,omitempty"`
	ScanType_Original                   string    `xml:"ScanType_Original" json:"scan_type_original,omitempty"`
	ScreenplayBy                        string    `xml:"ScreenplayBy" json:"screenplay_by,omitempty"`
	Season                              string    `xml:"Season" json:"season,omitempty"`
	Season_Position                     int       `xml:"Season_Position" json:"season_position,omitempty"`
	Season_Position_Total               int       `xml:"Season_Position_Total" json:"season_position_total,omitempty"`
	Service_Url                         string    `xml:"Service_Url" json:"service_url,omitempty"`
	ServiceChannel                      string    `xml:"ServiceChannel" json:"service_channel,omitempty"`
	ServiceKind                         string    `xml:"ServiceKind" json:"service_kind,omitempty"`
	ServiceName                         string    `xml:"ServiceName" json:"service_name,omitempty"`
	ServiceProvider                     string    `xml:"ServiceProvider" json:"service_provider,omitempty"`
	ServiceProviderr_Url                string    `xml:"ServiceProviderr_Url" json:"service_providerr_url,omitempty"`
	ServiceType                         string    `xml:"ServiceType" json:"service_type,omitempty"`
	SoundEngineer                       string    `xml:"SoundEngineer" json:"sound_engineer,omitempty"`
	Source_Duration                     float64   `xml:"Source_Duration" json:"source_duration,omitempty"`
	Source_Duration_FirstFrame          float64   `xml:"Source_Duration_FirstFrame" json:"source_duration_first_frame,omitempty"`
	Source_Duration_LastFrame           float64   `xml:"Source_Duration_LastFrame" json:"source_duration_last_frame,omitempty"`
	Source_FrameCount                   int64       `xml:"Source_FrameCount" json:"source_frame_count,omitempty"`
	Source_SamplingCount                int       `xml:"Source_SamplingCount" json:"source_sampling_count,omitempty"`
	Source_StreamSize                   int64       `xml:"Source_StreamSize" json:"source_stream_size,omitempty"`
	Source_StreamSize_Encoded           int64       `xml:"Source_StreamSize_Encoded" json:"source_stream_size_encoded,omitempty"`
	Source_StreamSize_Proportion        string    `xml:"Source_StreamSize_Proportion" json:"source_stream_size_proportion,omitempty"`
	Standard                            string    `xml:"Standard" json:"standard,omitempty"`
	Stored_Height                       int       `xml:"Stored_Height" json:"stored_height,omitempty"`
	Stored_Width                        int       `xml:"Stored_Width" json:"stored_width,omitempty"`
	StreamOrder                         string    `xml:"StreamOrder" json:"stream_order,omitempty"`
	StreamSize                          int64       `xml:"StreamSize" json:"stream_size,omitempty"`
	StreamSize_Encoded                  int64       `xml:"StreamSize_Encoded" json:"stream_size_encoded,omitempty"`
	StreamSize_Proportion               string    `xml:"StreamSize_Proportion" json:"stream_size_proportion,omitempty"`
	Subject                             string    `xml:"Subject" json:"subject,omitempty"`
	SubTrack                            string    `xml:"SubTrack" json:"sub_track,omitempty"`
	Summary                             string    `xml:"Summary" json:"summary,omitempty"`
	Synopsis                            string    `xml:"Synopsis" json:"synopsis,omitempty"`
	Tagged_Application                  string    `xml:"Tagged_Application" json:"tagged_application,omitempty"`
	Tagged_Date                         string    `xml:"Tagged_Date" json:"tagged_date,omitempty"`
	TermsOfUse                          string    `xml:"TermsOfUse" json:"terms_of_use,omitempty"`
	TextCount                           int       `xml:"TextCount" json:"text_count,omitempty"`
	ThanksTo                            string    `xml:"ThanksTo" json:"thanks_to,omitempty"`
	TimeCode_FirstFrame                 string    `xml:"TimeCode_FirstFrame" json:"time_code_first_frame,omitempty"`
	TimeCode_Settings                   string    `xml:"TimeCode_Settings" json:"time_code_settings,omitempty"`
	TimeCode_Source                     string    `xml:"TimeCode_Source" json:"time_code_source,omitempty"`
	TimeCode_Striped                    string    `xml:"TimeCode_Striped" json:"time_code_striped,omitempty"`
	TimeStamp_FirstFrame                float64   `xml:"TimeStamp_FirstFrame" json:"time_stamp_first_frame,omitempty"`
	TimeZone                            string    `xml:"TimeZone" json:"time_zone,omitempty"`
	TimeZones                           string    `xml:"TimeZones" json:"time_zones,omitempty"`
	Title                               string    `xml:"Title" json:"title,omitempty"`
	Title_More                          string    `xml:"Title_More" json:"title_more,omitempty"`
	Track                               string    `xml:"Track" json:"track,omitempty"`
	Track_More                          string    `xml:"Track_More" json:"track_more,omitempty"`
	Track_Position                      int64       `xml:"Track_Position" json:"track_position,omitempty"`
	Track_Position_Total                int64       `xml:"Track_Position_Total" json:"track_position_total,omitempty"`
	Track_Sort                          string    `xml:"Track_Sort" json:"track_sort,omitempty"`
	Transfer_characteristics            string    `xml:"transfer_characteristics" json:"transfer_characteristics,omitempty"`
	Transfer_characteristics_Original   string    `xml:"transfer_characteristics_Original" json:"transfer_characteristics_original,omitempty"`
	Transfer_characteristics_Source     string    `xml:"transfer_characteristics_Source" json:"transfer_characteristics_source,omitempty"`
	Type                                TrackType `xml:"type,attr" json:"type,omitempty"`
	UniqueID                            string    `xml:"UniqueID" json:"unique_id,omitempty"`
	UniversalAdID_Registry              string    `xml:"UniversalAdID_Registry" json:"universal_ad_id_registry,omitempty"`
	UniversalAdID_Value                 string    `xml:"UniversalAdID_Value" json:"universal_ad_id_value,omitempty"`
	VideoCount                          int       `xml:"VideoCount" json:"video_count,omitempty"`
	Width                               int       `xml:"Width" json:"width,omitempty"`
	Width_CleanAperture                 int       `xml:"Width_CleanAperture" json:"width_clean_aperture,omitempty"`
	Width_Offset                        int       `xml:"Width_Offset" json:"width_offset,omitempty"`
	Width_Original                      int       `xml:"Width_Original" json:"width_original,omitempty"`
	Written_Date                        string    `xml:"Written_Date" json:"written_date,omitempty"`
	Written_Location                    string    `xml:"Written_Location" json:"written_location,omitempty"`
	WrittenBy                           string    `xml:"WrittenBy" json:"written_by,omitempty"`
}

//GetFirstVideoTrack provides access to first Video track of media
//or nil if no video tracks found
func (m *MediaInfo) GetFirstVideoTrack() *Track {
	for _, t := range m.Media.Tracks {
		if t == nil {
			return nil
		}

		if t.Type == TrackVideo {
			return t
		}
	}

	return nil
}

//GetFirstAudioTrack provides access to first Audio track of media
//or nil if no audio tracks found
func (m *MediaInfo) GetFirstAudioTrack() *Track {
	for _, t := range m.Media.Tracks {
		if t == nil {
			return nil
		}

		if t.Type == TrackAudio {
			return t
		}
	}

	return nil
}

//GetTracksByType provides access to an array of tracks filtered by specific type
func (m *MediaInfo) GetTracksByType(trackType TrackType) (tracks []*Track) {
	for _, t := range m.Media.Tracks {
		if t == nil {
			return nil
		}

		if t.Type == trackType {
			tracks = append(tracks, t)
		}
	}

	return tracks
}
