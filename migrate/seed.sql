-- Clear existing data
DELETE FROM targeting_rules;
DELETE FROM campaigns;

-- Insert campaigns (no manual ID, only 'code')
INSERT INTO campaigns (code, name, image_url, cta, status)
VALUES 
  ('spotify', 'Spotify - Music for Everyone', 'https://upload.wikimedia.org/spotify.png', 'Download', 'ACTIVE'),
  ('duolingo', 'Duolingo - Learn Languages', 'https://upload.wikimedia.org/duolingo.png', 'Install', 'ACTIVE'),
  ('subwaysurfer', 'Subway Surfer', 'https://upload.wikimedia.org/subwaysurfer.png', 'Play', 'ACTIVE');

-- Add targeting rules using subquery to fetch campaign_id by code
-- Spotify
INSERT INTO targeting_rules (campaign_id, dimension, type, value)
SELECT id::BIGINT, 'country', 'include', 'us' FROM campaigns WHERE code = 'spotify';

INSERT INTO targeting_rules (campaign_id, dimension, type, value)
SELECT id::BIGINT, 'country', 'include', 'canada' FROM campaigns WHERE code = 'spotify';

-- Duolingo
INSERT INTO targeting_rules (campaign_id, dimension, type, value)
SELECT id::BIGINT, 'os', 'include', 'android' FROM campaigns WHERE code = 'duolingo';

INSERT INTO targeting_rules (campaign_id, dimension, type, value)
SELECT id::BIGINT, 'os', 'include', 'ios' FROM campaigns WHERE code = 'duolingo';

INSERT INTO targeting_rules (campaign_id, dimension, type, value)
SELECT id::BIGINT, 'country', 'exclude', 'us' FROM campaigns WHERE code = 'duolingo';

-- Subway Surfer
INSERT INTO targeting_rules (campaign_id, dimension, type, value)
SELECT id::BIGINT, 'os', 'include', 'android' FROM campaigns WHERE code = 'subwaysurfer';

INSERT INTO targeting_rules (campaign_id, dimension, type, value)
SELECT id::BIGINT, 'app', 'include', 'com.gametion.ludokinggame' FROM campaigns WHERE code = 'subwaysurfer';
