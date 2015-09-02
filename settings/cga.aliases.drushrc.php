<?php
/*
 * All you need to do is include an alias found
 * in local.aliases.drushrc.php as the 'parent'
 * for this environment and it will inherit those
 * settings.
 */
$aliases['local'] = array(
  'parent' => '@local.cga',
);

$aliases['dev'] = array(
  'parent' => '@local.defaults',
  'uri' => 'cga.e3develop.com',
  'root' => '/var/www/cga.e3develop.com/htdocs',
  'remote-host' => '162.249.104.85',
  'remote-user' => '401e',
  'path-aliases' => array(
    '%files' => 'sites/default/files',
    '%dump-dir' => '/home/401e/tmp'
  ),
);
