INSERT INTO partners (
    uid,
    name,
    endpoint,
    is_enabled,
    countries,
    device_types,
    min_bid_floor,
    blocked_categories
)
VALUES
(
    'dsp-alpha',
    'DSP Alpha',
    'http://mock-dsp:9001/bid',
    true,
    ARRAY['RU', 'KZ'],
    ARRAY['mobile', 'desktop'],
    0.5,
    ARRAY['gambling']
),
(
    'dsp-beta',
    'DSP Beta',
    'http://mock-dsp:9001/bid',
    true,
    ARRAY['US', 'DE'],
    ARRAY['desktop'],
    1.0,
    ARRAY[]::TEXT[]
),
(
    'dsp-gamma',
    'DSP Gamma',
    'http://mock-dsp:9001/bid',
    true,
    ARRAY[]::TEXT[],
    ARRAY[]::TEXT[],
    0,
    ARRAY['adult']
),
(
    'dsp-disabled',
    'Disabled DSP',
    'http://mock-dsp:9001/bid',
    false,
    ARRAY[]::TEXT[],
    ARRAY[]::TEXT[],
    0,
    ARRAY[]::TEXT[]
);